test:
	@go test -count=1 -vet=all -cover -coverprofile=set.out -covermode=atomic -race

.PHONY: lint

lint:
	@go run honnef.co/go/tools/cmd/staticcheck  -show-ignored -f stylish -checks all

.PHONY: build

check: lint test
