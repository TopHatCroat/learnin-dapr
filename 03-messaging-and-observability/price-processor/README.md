# Price Processor

This is an example of loading a message from pub sub, parsing it and storing the result into a state store.

### Run Node service with Dapr

1. Run the Go service app with Dapr:

```bash
PUB_SUB_NAME=price_pub_sub \
TOPIC_NAME=price \
dapr run --app-id price-processor --components-path ../deploy -- yarn start
```

### Clean up
```bash
dapr stop --app-id price-processor
```
