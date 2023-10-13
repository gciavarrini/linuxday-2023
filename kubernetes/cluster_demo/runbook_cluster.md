# Cluster Runbook

## Prerequisites

Install:

- `docker` or `podman`
- `minikube`
- `kubectl` or `oc`

## Start minikube cluster

```bash
minikube start -p minikube-intro-cluster
```

or

```bash
minikube profile list
```

```bash
minikube profile minikube-intro-cluster
```

## Open `minikube` dashboard

```bash { background=true }
minikube dashboard
```

## Check cluster status

### Pods

```sh { terminalRows=3 }
kubectl get pods
```

### Services

```bash { terminalRows=3 }
kubectl get services
```

or the _shortcut_

```bash { terminalRows=3 }
kubectl get svc
```

## Create objects

### Pod

```bash { closeTerminalOnSuccess=true interactive=false mimeType=text/x-yaml }
cat pod-nginx.yml
```

```bash { terminalRows=3 }
kubectl apply -f pod-nginx.yml
```

#### Verify status

```bash { terminalRows=3 }
kubectl get pods
```

> BONUS  
> Use `kubectl get pods -w` to watch for changes

If pods status isn't `Running` (yet), let's wait

```bash
kubectl wait --for=condition=Ready pod/my-nginx --timeout=300s
```

### Service

```bash { closeTerminalOnSuccess=true interactive=false mimeType=text/x-yaml }
cat service.yml
```

```bash { terminalRows=3 }
kubectl apply -f service.yml
```

#### Verify status

```bash { terminalRows=4 }
kubectl get services
```

### Access `nginx`

```bash
minikube service my-service
```

## Deployments

### Create a ConfigMap

Let's customize the `ngnix` server

> NOTE: There are two config maps. Each config map has a different HTML page.

```bash { closeTerminalOnSuccess=true interactive=false mimeType=text/x-yaml }
cat config_maps.yml
```

```bash { terminalRows=2 }
kubectl apply -f config_maps.yml
```

#### Verify status

```bash { terminalRows=4 }
kubectl get cm
```

### Create deployments

```bash { terminalRows=4 }
kubectl get deployments
```

```bash { terminalRows=7 }
kubectl get pods
```

> NOTE: There are two deployments. Each deployment has 2 replicas

```bash { closeTerminalOnSuccess=true interactive=false mimeType=text/x-yaml }
cat deployments.yml
```

```bash { terminalRows=2 }
kubectl apply -f deployments.yml
```

Verify again deployments and pods.

### Service

```bash { closeTerminalOnSuccess=true interactive=false mimeType=text/x-yaml }
cat service-deployments.yml
```

```bash { terminalRows=3 }
kubectl apply -f service-deployments.yml
```

```bash { terminalRows=4 }
kubectl get services
```

### Access `nginx`

```bash
minikube service nginx-service-deployments
```

```bash { interpreter= mimeType=text/plain terminalRows=50 }
# Define the URL to be invoked
URL="$(minikube service nginx-service-deployments --url)"

# Number of requests to make
REQUESTS=10

# Loop to make HTTP requests
for ((i=1; i<=$REQUESTS; i++)); do
  echo "Request $i: \n $(lynx -dump "$URL" )\n"
done
```

# Stop minikube

```bash
minikube stop
```

or

```bash
minikube delete
```
