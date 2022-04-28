# Coindesk Fetcher

This is an example of fetching a resource from an API and publishing the parsed result on a Redis pub sub event queue.

### Run Go service with Dapr

1. Build the app with: `go build app.go`

2. Run the Go service app with Dapr:

```bash
COIN_DESK_ENDPOINT=https://api.coindesk.com/v1/bpi/currentprice.json \
PUB_SUB_NAME=price_pub_sub \
TOPIC_NAME=price \
dapr run --app-id coindesk-fetcher --components-path ../deploy-local -- go run main.go
```

### Clean up
```bash
dapr stop --app-id coindesk-fetcher
```
