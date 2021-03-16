# Builder
FROM golang:1.16-alpine AS builder

RUN apk --no-cache add git g++ linux-headers

COPY ./ /go/src
WORKDIR /go/src/
RUN GOOS=linux GOARCH=amd64 go build -o /main .

# Service
FROM alpine:latest AS runtime

RUN apk --no-cache add ca-certificates

WORKDIR /root
COPY --from=builder /main .
USER root
CMD ["./main"]

EXPOSE 9153
