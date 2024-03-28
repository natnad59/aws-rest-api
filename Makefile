build:
	go build -o out/api

clean:
	rm -rf out/* && \
	rm -rf lambda/*

push:
	env GOOS=linux GOARCH=amd64 go build -o lambda/bootstrap && \
	cd lambda && \
	zip -r rest-api-example.zip bootstrap && \
	aws s3 cp rest-api-example.zip s3://backend-dev-artifacts/examples/rest-api-example.zip

apply:
	cd terraform && \
	terraform init && \
	terraform workspace select dev || terraform workspace new dev && \
	terraform apply -auto-approve

destroy:
	cd terraform && \
	terraform init && \
	terraform workspace select dev || terraform workspace new dev && \
	terraform destroy -auto-approve

deploy:
	cd scripts && \
	./deploy.sh
