# State and k8s

This example will have show how to handle state with Kubernetes and Redis.
The service will offer 2 simple methods:
* `POST /increment` - increments the counter
* `GET /counter` - show the current counter state

### Local setup

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


### Kubernetes setup

For Kubernetes setup first we need to make sure there is a local Minikube cluster runnin:
`minikube start`

Then initialise dapr within the cluster with `dapr init --kubernetes`

Running `kubectl get pods -A` should now show dapr related pods running

### Building and running the images

For these images to be available to your minikube cluster, in your terminal,
before building them you must run: `eval $(minikube -p minikube docker-env)`

To build both containers run:
```bash
(cd ./counter && docker build -t counter-app:latest .)
(cd ./multiplier && docker build -t multiplier-app:latest .)
```

Now we have two containers ready to deploy to k8s, but first, we need to deploy Redis:
`kubectl apply -f ./k8s/redis.yml`

Then we can run our two containers with: `kubectl apply -f ./k8s/app.yml`

### Clean up 

Remove created containers with:
```bash
kubectl delete -f ./k8s/app.yml
kubectl delete -f ./k8s/redis.yml
```
