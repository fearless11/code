FROM golang:1.18 as builder

ENV GOPROXY https://proxy.golang.com.cn,direct
WORKDIR /app
COPY . /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main -v

FROM alpine:3.17.0
EXPOSE 8080
WORKDIR /app
RUN apk update
RUN apk add curl busybox-extras
COPY --from=builder /app /app
CMD ["/app/main"]
