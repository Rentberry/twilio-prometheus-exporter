FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root
COPY dist/twilio-prometheus-exporter_linux_amd64/twilio-prometheus-exporter .
USER root
CMD ["./twilio-prometheus-exporter"]

EXPOSE 9153
