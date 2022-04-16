# State and k8s

This example will have show how to handle state with Kubernetes and Redis.
The service will offer 2 simple methods:
* `POST /increment` - increments the counter
* `GET /counter` - show the current counter state

### Bare startup

Since dapr comes with a small Redis instance by default, we can just run this app alone:
`dapr run --app-id counter-app --app-port 8088 --dapr-http-port 8089 go run main.go`

### Invoking the app

You can increment the counter with: `dapr invoke --verb POST --app-id counter-app --method increment`

Or with curl: `curl -X POST http://localhost:8089/v1.0/invoke/counter-app/method/increment`

And get the state of the counter with: `dapr invoke --verb GET --app-id counter-app --method counter`
or: `curl http://localhost:8089/v1.0/invoke/counter-app/method/counter`

Notice that if you kill the app with `Ctrl-C` and start it again, the state of the counter will
persist since the Redis store is an external service in this case. 

### Clean up

And finally, stop the app with: `dapr stop --app-id counter-app`


