build:
	protoc -I. --go_out=plugins=grpc:. proto/consignment/consignment.proto
	
build-docker:
	GOOS=linux GOARCH=amd64 go build
	docker build -t kaansari/consignment:latest .
	docker push kaansari/consignment:latest


run:
	docker run -p 50051:50051 \
		-e DB_HOST=localhost \
		-e DB_PASS=password \
		-e DB_USER=postgres \
		kaansari/consignment

deploy:
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yml
	kubectl replace -f ./deployments/deployment.yml
