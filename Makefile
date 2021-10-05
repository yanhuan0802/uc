IMAGE_NAME=yanhuan6252/httpserver
IMAGE_VERSION=1.0.0

buildAndPush:
	docker buildx build --platform linux/amd64,linux/arm64  -t $(IMAGE_NAME):$(IMAGE_VERSION) --push .