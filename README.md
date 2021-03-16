# Twilio usage exporter for Prometheus
```
docker run -p 9860:9860 \
    -e TWILIO_EXPORTER_ACCOUNT_ID=your-twilio-account-id \
    -e TWILIO_EXPORTER_SID=your-twilio-sid \
    -e TWILIO_EXPORTER_API_KEY=your-twilio-api-key \
    rentberry/twilio-prometheus-exporter:latest
```

## What's exported?
- ``twilio_count`` - count of usage events per account and category
- ``twilio_usage`` - the amount used to bill usage and measured in usage_units
- ``twilio_price`` - the total price of the usage in the currency specified in price_unit and associated with the account.
