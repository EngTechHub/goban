# Go parameters
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOTOOL=$(GOCMD) tool
BINARY_NAME=mybinary
BINARY_UNIX=$(BINARY_NAME)_unix


cover:
        $(GOTEST) -coverprofile=covprofile
        $(GOTOOL) COVER -HTML=covprofile -o coverage.html