FROM golang:1.10.0-alpine3.7 as builder
RUN apk add --update --no-cache alpine-sdk ca-certificates \
      libressl \
      git openssh openssl build-base coreutils upx
WORKDIR /go/src/github.com/StevenACoffman/logsink/dummy-logsink
RUN go get -d -v github.com/apex/log
RUN go get -d -v github.com/gorilla/pat
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -s' -o main main.go
RUN upx --brute main

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/StevenACoffman/logsink/dummy-logsink/main /
CMD ["/main"]
