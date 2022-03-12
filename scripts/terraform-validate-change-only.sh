#!/usr/bin/env bash

set -e

for f in terraform/*; do
    if [ -d $f ]; then
        cd $f/prod
        CHANGES=$(git diff --name-only HEAD..origin/main ..)
        # if there are any changes run terraform apply
        if [ "$CHANGES" != "" ]; then
            echo "detected changes in $f. Running terraform apply..."
            terraform init -no-color
            terraform validate -no-color
        fi
        cd -
    fi
done
