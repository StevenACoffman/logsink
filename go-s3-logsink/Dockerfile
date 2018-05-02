FROM golang:1.10.0-alpine3.7 as builder
RUN apk add --update --no-cache alpine-sdk ca-certificates \
      libressl \
      git openssh openssl build-base coreutils upx
WORKDIR /go/src/github.com/StevenACoffman/logsink/go-s3-logsink
RUN go get -d -v github.com/apex/log
RUN go get -d -v github.com/aws/aws-sdk-go
RUN go get -d -v github.com/satori/go.uuid
RUN go get -d -v github.com/gorilla/pat
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -s' -o main main.go
RUN upx --brute main

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENV AWS_DEFAULT_REGION "us-east-1"
ENV AWS_REGION "us-east-1"
ENV AWS_DEFAULT_OUTPUT "json"
COPY --from=builder /go/src/github.com/StevenACoffman/logsink/go-s3-logsink/main /
CMD ["/main"]
