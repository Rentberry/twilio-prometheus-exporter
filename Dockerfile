FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY ./twilio-prometheus-exporter .
ENTRYPOINT ["./twilio-prometheus-exporter"]

EXPOSE 9860
