# parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test 
GOCOVERAGE=$(GOCMD) tool cover
BINARY_NAME=batch-funds-loader
SRC_PATH=cmd/$(BINARY_NAME)/main.go

all: test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) $(SRC_PATH)
test: 
	$(GOTEST) -v ./...
coverage:
	$(GOTEST) ./... -coverprofile cp.out
	$(GOCOVERAGE) -html=cp.out
clean: 
	$(GOCLEAN) $(SRC_PATH)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
