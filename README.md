# Exploring kubernetes

## Installations

### kind

Kind is used to create kubernetes cluster 

```bash
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.15.0/kind-linux-amd64
chmod +x ./kind
sudo mv ./kind /usr/local/bin/kind
```


## Dummy server

1. Build docker
    ```
    cd dummy_server/
    docker build --tag dummy-server:v1 .
    ```
2. Docker run command
    ```
    docker run --rm --name server -p 8080:8080 dummy-server:v1
    ```


## Getting started

### Setup the kubernetes cluster

1. Create a kubernetes cluster
    ```bash
    kind create cluster
    kubectl cluster-info --context kind-kind  # Verify the cluster info
    ```
2. Load the docker images into the cluster
    ```bash
    kind load docker-image dummy-server:v1
    kind load docker-image dummy-server:v2
    ```
3. Get the images loaded in the node
    ```bash
    kubectl get nodes
    docker exec -ti <node-name> crictl images
    ```

### Deploy the app

```bash
kubectl create deployment dummy-server --image=dummy-server:v1
```

Other useful commands

```bash
kubectl get deployments  # Get deployment information
kubectl get pods  # Get pods information
export POD_NAME=$(kubectl get pods -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')  # Fetch pod name
kubectl describe pods  # Describes the pods
kubectl describe deployments  # Describes the deployments
kubectl logs $POD_NAME  # Get logs from the pod
kubectl exec $POD_NAME -- env  # Get environments of a pod
kubectl exec -ti $POD_NAME -- /bin/sh  # Enter into the pod container
```

### Expose the port of the app

```bash
kubectl expose deployment/kubernetes-bootcamp --type="NodePort" --port 8080
kubectl get services  # List out the services, when we expose a deployment or pod it creates a service
```

Verify the app

```bash
export NODE_PORT=$(kubectl get services/dummy-server -o go-template='{{(index .spec.ports 0).nodePort}}')  # Fetch the node port exposed for this app
kubectl get nodes -o wide  # Get the node ip from the output of this command
curl $NODE_IP:$NODE_PORT
```

### Scale the app

```bash
kubectl get deployments
kubectl get rs  # Get replica statuses
kubectl describe deployments/dummy-server
kubectl scale deployments/dummy-server --replicas=2  # Scale up to 2 replicas
kubectl describe deployments/dummy-server
```

### Rolling update

If we keep running a simple script which keeps fetching from the server, initially the response would be `This is v1` and it would change to `This is v2` as the update completes. This shows us two things, first, the update has taken place, and second, there is no **`downtime`**.

```bash
kubectl set image deployments/dummy-server dummy-server=dummy-server:v2  # Update the image and k8s will take care of rolling out the update without downtime
kubectl get pods -o wide
kubectl describe deployments/dummy-server
kubectl scale deployments/dummy-server --replicas=1  # Scale down to 1 replica
```

If the update had a bug/issue, we can rollback to previous version

```bash
kubectl rollout undo deployments/dummy-server
kubectl describe services
kubectl describe pods
kubectl describe deployment
```

### Delete the deployments

```bash
kubectl delete deployments/dummy-server
kubectl get deployments
kubectl get pods -o wide
```

Delete the `kind` cluster

```bash
kind delete cluster
kind get clusters
```


## References

1. [Kubernetes basics](https://kubernetes.io/docs/tutorials/kubernetes-basics/)
2. [Kubernetes cheatsheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)
