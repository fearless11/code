FROM golang:1.18-alpine 

ENV GOPROXY https://proxy.golang.com.cn,direct
WORKDIR /app
COPY . /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main -v

CMD ["/app/main"]
