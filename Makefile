PORT?=8080
IMAGE_NAME:=go-reverse-proxy
CONTAINER_NAME:=go-reverse-proxy
VOLUME_PATH:=/.entry #you probabily dont need this. But i uses multiple containers at my machine and need to pay attention with my disk space

build:
	docker build -t $(IMAGE_NAME) .

run: 
	docker run -d --name $(CONTAINER_NAME) -v $(VOLUME_PATH):/app -p $(PORT):$(PORT) $(IMAGE_NAME)

stop:
	docker stop $(CONTAINER_NAME) && \
	docker rm $(CONTAINER_NAME) &

clear:
	docker rmi $(IMAGE_NAME) && \
	docker rmi $(CONTAINER_NAME) &