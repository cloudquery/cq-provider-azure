package client

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var errorCodeDescriptions = map[interface{}]string{
	http.StatusNotFound:   "The requested resource could not be found.",
	http.StatusBadRequest: "Bad request",
	http.StatusForbidden:  "You are not authorized to perform this operation.",
}

func ErrorClassifier(meta schema.ClientMeta, resourceName string, err error) diag.Diagnostics {
	client := meta.(*Client)

	// Don't override if already a diagnostic, just redact
	if d, ok := err.(diag.Diagnostic); ok {
		return diag.Diagnostics{
			RedactError(client.SubscriptionId, d),
		}
	}

	var detailedError autorest.DetailedError
	if errors.As(err, &detailedError) {
		switch detailedError.StatusCode {
		case http.StatusNotFound:
			return diag.Diagnostics{
				RedactError(client.SubscriptionId, diag.NewBaseError(err, diag.IGNORE, diag.RESOLVING, resourceName, ParseSummaryMessage(client.SubscriptionId, err, detailedError), errorCodeDescriptions[detailedError.StatusCode])),
			}
		case http.StatusBadRequest:
			return diag.Diagnostics{
				RedactError(client.SubscriptionId, diag.NewBaseError(err, diag.WARNING, diag.RESOLVING, resourceName, ParseSummaryMessage(client.SubscriptionId, err, detailedError), errorCodeDescriptions[detailedError.StatusCode])),
			}
		case http.StatusForbidden:
			return diag.Diagnostics{
				RedactError(client.SubscriptionId, diag.NewBaseError(err, diag.WARNING, diag.ACCESS, resourceName, ParseSummaryMessage(client.SubscriptionId, err, detailedError), errorCodeDescriptions[detailedError.StatusCode])),
			}
		}
	}

	// Take over from SDK and always return diagnostics, redacting PII
	return diag.Diagnostics{
		RedactError(client.SubscriptionId, diag.NewBaseError(err, diag.ERROR, diag.RESOLVING, resourceName, err.Error(), "")),
	}
}

func ParseSummaryMessage(subscriptionId string, err error, detailedError autorest.DetailedError) string {
	for {
		if de, ok := err.(autorest.DetailedError); ok {
			return fmt.Sprintf("%s: %s - %s", de.Method, de.PackageType, detailedError.Error())
		}
		if err = errors.Unwrap(err); err == nil {
			return detailedError.Error()
		}
	}
}

// RedactError redacts a given diagnostic and returns a RedactedDiagnostic containing both original and redacted versions
func RedactError(subId string, e diag.Diagnostic) diag.Diagnostic {
	r := diag.NewBaseError(
		errors.New(removePII(subId, e.Error())),
		e.Severity(),
		e.Type(),
		e.Description().Resource,
		removePII(subId, e.Description().Summary),
		removePII(subId, e.Description().Detail),
	)
	return diag.NewRedactedDiagnostic(e, r)
}

var (
	uuidRegex = regexp.MustCompile(`(\W)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}(\W)`)
)

func removePII(subId string, msg string) string {
	msg = strings.ReplaceAll(msg, subId, "xxxx")
	msg = uuidRegex.ReplaceAllString(msg, "$1xxxx$2")
	return msg
}
