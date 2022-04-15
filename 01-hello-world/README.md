# Hello dapr

This is a dead simple example of running dapr.

First you need to get a release binary from:

### Start up

Then initialise dapr with: `dapr init`

To run this example use: `dapr run --app-id hello-dapr --app-port 8088 --dapr-http-port 8089 go run main.go`

### Invoking the app

Now invoke the app with: `dapr invoke --verb GET --app-id hello-dapr --method greeting`

Or with curl: `curl http://localhost:8089/v1.0/invoke/hello-dapr/method/greeting`

### Clean up

And finally, stop the app with: `dapr stop --app-id hello-dapr`