#!/bin/sh
gcloud auth activate-service-account --key-file=/secrets/skaffold-service-account.json
gcloud config set project $PROJECT_DEV_NAME
gcloud container clusters get-credentials $CLUSTER_DEV_NAME --region europe-west1-b
kubectl config get-contexts 
export STATE=$(git rev-list -1 HEAD --abbrev-commit)  
skaffold build -m $SERVICE_NAME -p prod --file-output build-$STATE.json
skaffold deploy -m $SERVICE_NAME -a build-$STATE.json