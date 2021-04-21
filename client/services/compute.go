package services

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/profiles/2020-09-01/compute/mgmt/compute"
	"github.com/Azure/go-autorest/autorest"
)

type ComputeClient struct {
	Disks DisksClient
}

func NewComputeClient(subscriptionId string, auth autorest.Authorizer) ComputeClient {
	disks := compute.NewDisksClient(subscriptionId)
	disks.Authorizer = auth
	return ComputeClient{
		Disks: disks,
	}
}

type DisksClient interface {
	List(ctx context.Context) (result compute.DiskListPage, err error)
}
