FROM golang:1.10.0-alpine3.7 as builder
RUN apk add --update --no-cache alpine-sdk ca-certificates \
      libressl \
      git openssh openssl build-base coreutils
WORKDIR /go/src/github.com/StevenACoffman/logsink/dummy-logsink
RUN go get -d -v github.com/apex/log
RUN go get -d -v github.com/gorilla/pat
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM golang:1.10.0-alpine3.7
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/github.com/StevenACoffman/logsink/dummy-logsink/app .
CMD ["./app"]
