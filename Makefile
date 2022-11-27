GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_LOC=bin
BINARY_NAME=trokhos
DOCKER_REPOSITORY_OWNER=alwindoss
VERSION=0.0.1

all: build
docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./$(BINARY_LOC)/$(BINARY_NAME) -v ./cmd/$(BINARY_NAME)/...
package:
	docker build -t $(DOCKER_REPOSITORY_OWNER)/$(BINARY_NAME):$(VERSION) .
publish:
	docker push $(DOCKER_REPOSITORY_OWNER)/$(BINARY_NAME):$(VERSION)
setup:
	$(GOGET) -v ./...
build:
ifeq ($(OS),Windows_NT)
	$(GOBUILD) -o ./$(BINARY_LOC)/$(BINARY_NAME).exe -v ./cmd/$(BINARY_NAME)/...
else
	$(GOBUILD) -o ./$(BINARY_LOC)/$(BINARY_NAME) -v ./cmd/$(BINARY_NAME)/...
endif 
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -rf $(BINARY_LOC)
run: clean build
ifeq ($(OS),Windows_NT)
	./$(BINARY_LOC)/$(BINARY_NAME).exe
else
	./$(BINARY_LOC)/$(BINARY_NAME)
endif 