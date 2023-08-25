test:
	@go test -count=1 -vet=all -cover -coverprofile=set.out -covermode=atomic -race
