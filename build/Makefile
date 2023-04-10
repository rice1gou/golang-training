.PHONY: dev
start:
	GO_MOD=backend docker-compose up
start-build:
	GO_MOD=backend docker-compose up --build
cleanup:
	docker-compose down --rmi all --volumes --remove-orphans
destroy:
	docker-compose down --volumes --remove-orphans

.PHONY: test
test:
	go test
lint:
	golangci-lint run -c ./scripts/configs/.golangci.yaml
