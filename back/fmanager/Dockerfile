#build stage
FROM golang:alpine AS builder
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories  \ 
    && apk update && apk add --no-cache git && go env -w  GOPROXY=https://goproxy.cn,direct \
    && mkdir -p /go/bin/app
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
COPY --from=builder /go/src/app/config.yaml  /app
COPY --from=builder /go/src/app/dist  /app/dist
ENV GIN_MODE=release
ENTRYPOINT /app/fmanager
LABEL Name=fmanager Version=1.0
EXPOSE 3000
