IMAGE_NAME=yanhuan6252/httpserver
IMAGE_VERSION=1.0.0

build:
	docker build -t $(IMAGE_NAME):$(IMAGE_VERSION) .

push: build
	docker push $(IMAGE_NAME):$(IMAGE_VERSION)

multiBuildAndPush:
	docker buildx build --platform linux/amd64,linux/arm64  -t $(IMAGE_NAME):$(IMAGE_VERSION) --push .

dockerRun:
	docker run -d --name httpserver -p 80:80 $(IMAGE_NAME):$(IMAGE_VERSION)