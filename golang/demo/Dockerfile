FROM golang:1.18 as builder

ENV GOPROXY https://proxy.golang.com.cn,direct
WORKDIR /app
COPY . /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main -v

FROM 11expose11/alpine:3.17.1
EXPOSE 8080
WORKDIR /app
COPY --from=builder /app /app
CMD ["/app/main"]
