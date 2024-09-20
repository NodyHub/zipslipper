default: build

build:
	@go build -o zipslipper .
	

install: build
	@mv zipslipper $(GOPATH)/bin/zipslipper

clean:
	@go clean
	@rm zipslipper

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

test_coverage_view:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

test_coverage_html:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o=coverage.html

all: build install