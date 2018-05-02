<img height="100" src="./images/logsink.svg" />

# logsink

The purpose of this repository is to recieve batched log messages (from fluent-bit) and ship them to AWS S3 and Apache Kafka.

## go-kafka-logsink

This container will recieve batched log messages from fluent-bit (inside kubernetes) and send them to the proper Kafka topics. One message for on each topic for each Dests for Ithaka structured log messages, or whatever topic was annotated application default topic for unstructured messages

## go-s3-logsink

This container will recieve batched log messages from fluent-bit (inside kubernetes) and send them to an S3 bucket. The output format is intentionally identical to the input signature of go-kafka-logsink for easy re-ingestion in the event of kafka topic deletion.

## dummy-logsink

The Dummy-Logsink container receives a batch of raw log messages from fluent-bit (including metadata) via HTTP Post and responds with a status of 200 OK. Useful for reproducing issues.

## echo-server

The echo-server container recieves a batch of raw log messages from fluent-bit (including metadata) via HTTP Post and writes them out to standard out for debugging purposes.