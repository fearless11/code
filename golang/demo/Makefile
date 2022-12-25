APP=demo
VERSION=v0.1

all: test build-linux

run:
	@echo "> go run"
	go run main.go

test:
	@echo "> go test"
	go test -v ./...

clean:
	@echo "> go clean"
	go clean && rm -f $(APP)

build-mac:
	@echo "> go build"
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(APP) -v 

build-linux:
	@echo "> go build"
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(APP) -v 

docker-build:
	@echo "> docker build"
	docker build -t $(APP):$(VERSION) .

docker-run:
	@echo "> docker run"
	docker run -d \
	--name=${APP} \
	--net=host \
    -m 256M \
    --restart=always \
    --log-opt max-size=100m \
    --log-opt max-file=10 \
    $(APP):$(VERSION)

docker-stop:
	@echo "> docker stop"
	docker stop ${APP}
	docker rm ${APP}

docker-exec:
	@echo "> docker exec"
	docker exec -it ${APP} sh

docker-logs:
	@echo "> docker logs"
	docker logs --tail=10 -f ${APP}
