# Welcome to Coffee Shop App!

This app built with golang and postgres database to check the invalid deliveries in coffee shop delivery request


## Installation

You must first to build docker image using the following steps
```
#building docker image with image take coffer-shop
docker build . -t coffee-app/coffee-shop:latest
#push image to container repository 
docker push coffee-app/coffee-shop:latest
```
Please note you must have access to docker hub to push this image for example if you are using dockerhub as container registry

## Configuration

the app can be deployed to kubernetes by applying the yaml files in yamls directory but you need first to change the database connection variables in **k8s-yamls/coffee-shop.deployment.yaml** the environment variables as the following
```
- name: DB_HOST
value: "Database Host"

- name: DB_USER
value: "Database Username"

- name: DB_PASSWORD
value: "Database User Password"

- name: DB_NAME
value: "Database Name"
```
Also you can change the app listening port by changing the following variable in file **k8s-yamls/coffee-shop.deployment.yaml**
```
- name: SERVER_PORT
value: "new server listening port"
```
but in this case you must set the following directives in **k8s-yamls/coffee-shop.deployment.yaml** with the same port you entered in **SERVER_PORT** variable

```
#line 30
containerPort: "new server listening port" 
```
and the file **k8s-yamls/coffee-shop.service.yaml** must be also changed
```
#line 7,9
port: "new server listening port"
targetPort: "new server listening port"
```

## Deploy to Kubernetes
kubectl create -f k8s-yamls/

## Functionality Verifying 
In case of we deployed the app in default namespace you can check the statues of pod 
```
kubectl get pods --selector=app=coffee-shop
```

Also you can test the service accessibility using port-forward
```
kubectl port-forward svc/coffee-shop 3000:3000
#please make sure your local port 3000 is available if not you can change the forwarding port to i.e 8080
kubectl port-forward svc/coffee-shop 8080:3000
``` 

Also you can check pod logs for debugging
```
kubectl logs --selector=app=coffee-shop
```

## APIs 
After you did the port-forward you can go through your web browser to check the following APIs endpoints

http://localhost:3000/api/v1/invalid-deliveries or http://localhost:8080/api/v1/invalid-deliveries the page must render json result for invalid deliveries also you can use curl to test this API endpoint
```
curl --silent http://localhost:3000/api/v1/invalid-deliveries
or
curl --silent http://localhost:8080/api/v1/invalid-deliveries
```
