APP_NAME=go_server
BINARY_NAME=go_server
VERSION=latest

GOOS=linux
GOARCH=amd64

build:
	@echo "🔨 Building..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINARY_NAME)

clean:
	@echo "🧹 Cleaning up..."
	rm -f $(BINARY_NAME)

docker-build: build
	@echo "🐳 Building Docker image..."
	docker build -t $(APP_NAME):$(VERSION) .

docker-run:
	@echo "🚀 Running container on port 8080..."
	docker run -p 8080:8080 $(APP_NAME):$(VERSION)

push:
	@echo "☁️  Pushing image..."
	docker tag $(APP_NAME):$(VERSION) your-docker-user/$(APP_NAME):$(VERSION)
	docker push your-docker-user/$(APP_NAME):$(VERSION)
