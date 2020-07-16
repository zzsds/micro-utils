GOPATH:=$(shell go env GOPATH)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_TAG=$(shell git describe --abbrev=0 --tags --always --match "v*")
# IMAGE_TAG=$(GIT_TAG)-$(GIT_COMMIT)
IMAGE_TAG = 0.0.1
NAME = micro
IMAGE_NAME = micro-welfare/micro
NETWORK = micro-welfare
MICRO_REGISTER = mdns
TYPE = api
HANDLER = api
PORT = 8080

api:
	go run main.go --cors-allowed-headers="Content-Type,Authorization,Client-Type" --cors-allowed-origins="*" --cors-allowed-methods="OPTIONS,DELETE,GET,POST,PUT" --enable_stats api --handler=${HANDLER} --address=:${PORT}

web:
	go run main.go --enable_stats web 

build:
	go build -o ${NAME} *.go

test:
	go test -v ./... -cover



docker: docker-build micro-api

docker-build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(IMAGE_NAME):latest
	# docker push $(IMAGE_NAME):$(IMAGE_TAG)
	# docker push $(IMAGE_NAME):latest

micro-api:
	docker run --rm --name ${NAME}-api -d -p ${PORT}:${PORT} --network ${NETWORK} -e MICRO_REGISTER=${MICRO_REGISTER} ${IMAGE_NAME}:${IMAGE_TAG} --cors-allowed-headers="Content-Type,Authorization,Client-Type" --cors-allowed-origins="*" --cors-allowed-methods="OPTIONS,DELETE,GET,POST,PUT" --enable_stats api --handler=${HANDLER} --address=:${PORT}

micro-web:
	docker run --rm --name ${NAME}-web -d -p ${PORT}:${PORT} --network ${NETWORK} -e MICRO_REGISTER=${MICRO_REGISTER} ${IMAGE_NAME}:${IMAGE_TAG} --cors-allowed-headers="Content-Type,Authorization,Client-Type" --cors-allowed-origins="*" --cors-allowed-methods="OPTIONS,DELETE,GET,POST,PUT" --enable_stats web  --address=:${PORT}

PHONY: api web build docker-build micro-api micro-web 
