tools-install:
	go list -e -f '{{ join .Imports "\n" }}' -tags tools ./tools | xargs -I {} go install {}

tidy:
	go mod tidy
	go mod vendor

generate:
	go generate ./internal/...

unit-test:
	ginkgo -covermode atomic ./internal/...
