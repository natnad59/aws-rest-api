#!/bin/bash

set -e

pushd ..
env GOOS=linux GOARCH=amd64 go build -o lambda/bootstrap
pushd lambda
zip -r rest-api-example.zip bootstrap
ARTIFACTS_BUCKET=$(aws ssm get-parameter --name "/rest-api-example/artifacts-bucket" --query "Parameter.Value" --with-decryption --output text)
aws s3 cp rest-api-example.zip s3://${ARTIFACTS_BUCKET}/examples/rest-api-example.zip
popd

pushd terraform
terraform init
terraform workspace select dev || terraform workspace new dev
terraform apply -auto-approve
sleep 5
aws lambda update-function-code --function-name rest-api-example --s3-bucket ${ARTIFACTS_BUCKET} --s3-key examples/rest-api-example.zip
popd

sleep 5

# update api gateway with new lambda code
REST_API=$(aws ssm get-parameter --name "/rest-api-example/api-gateway/id" --query "Parameter.Value" --with-decryption --output text)
OLD_DEPLOYMENT=$(aws apigateway get-stages --rest-api-id ${REST_API} --query 'item[?stageName==`v1`].deploymentId' --output text)
aws apigateway update-stage --rest-api-id ${REST_API} --stage-name "v1" --patch-operations op=replace,path=/deploymentId,value=$(aws apigateway create-deployment --rest-api-id ${REST_API} --query id --output text)
aws apigateway delete-deployment --rest-api-id ${REST_API} --deployment-id ${OLD_DEPLOYMENT}
