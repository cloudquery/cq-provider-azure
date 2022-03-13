#!/usr/bin/env bash

# The infracost logic is from here # https://www.infracost.io/docs/troubleshooting/#multi-projects

set -e

plans=()
planjsons=()
for f in terraform/*; do
    if [ -d $f ]; then
        CHANGES=$(git diff --name-only HEAD..origin/main ${f})
        # if there are any changes run terraform apply
        if [ "$CHANGES" != "" ]; then
            echo "detected changes in $f. Running terraform plan..."
            terraform plan -chdir="${f}/prod" -lock=false -no-color -out tfplan.binary
            
            echo "Running terraform show in $f"
            terraform show -chdir="${f}/prod" -json tfplan.binary > "${f}/prod/plan.json"
            plans=(${plans[@]} "${f}/prod/tfplan.binary")
            planjsons=(${planjsons[@]} "${f}/prod/plan.json")
        fi
    fi
done

echo -e "version: 0.1\n\nprojects:\n" > /tmp/infracost-generated.yml
for planjson in "${planjsons[@]}"; do
  echo -e "  - path: $planjson" >> /tmp/infracost-generated.yml
done