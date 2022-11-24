APP=webhook
VERSION=v0.1

docker-build:
	@echo "> docker build"
	docker build -t $(APP):$(VERSION) .

docker-run:
	@echo "> docker run"
	docker run -d \
	--name=${APP} \
	--net=host  \
    -m 256M \
    --restart=always \
    --log-opt max-size=100m \
    --log-opt max-file=10 \
    $(APP):$(VERSION) \
    /app/main

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
