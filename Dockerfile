FROM golang:alpine

ENV GOPROXY "https://goproxy.cn,direct"

COPY . /workdir
WORKDIR /workdir

RUN go mod vendor
RUN go build -mod=vendor -o main main.go

FROM alpine

COPY --from=0 /workdir/main /workdir/
WORKDIR /workdir
ENV TZ Asia/Shanghai
RUN apk add bash

ENTRYPOINT ["./main"]