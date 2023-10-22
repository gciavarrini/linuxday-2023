# SonataFlow Demo

## Prerequisites

Install:

- `docker` or `podman`
- `minikube`
- `kn`
- `kn-workflow`
- `k9s` (optional)

## Create Minikube cluster

```bash
minikube start --cpus 8 --memory 8192 --addons registry --addons metrics-server --insecure-registry "10.0.0.0/24" --insecure-registry "localhost:5000"
```

or

```bash { terminalRows=3 }
minikube profile minikube-sonata
```

## Verify

```bash { terminalRows=5 }
minikube profile list 
```

## Create the tunnel

```bash
minikube tunnel --profile minikube-sonata
```

## Create the SonataFlow Operator

```bash
kubectl create -f https://raw.githubusercontent.com/kiegroup/kogito-serverless-operator/v1.43.0/operator.yaml
```

### Wait

```bash
kubectl get pod -n sonataflow-operator-system --watch
```

### Access operator logs

```bash
 kubectl logs deployment/sonataflow-operator-controller-manager -n sonataflow-operator-system -f
```

## Create the workflow

```bash
kn workflow create --name serverless-hello-world
kn workflow deploy
```

```bash { closeTerminalOnSuccess=true cwd=serverless-hello-world interactive=false mimeType=text/x-json }
code workflow.sw.json
```

## Check the workflow

```bash { terminalRows=4 }
kubectl get workflow
```

```bash
kubectl get service
```

> HINT: use `k9s` to inspect the objects

### Get the service URL

```bash { terminalRows= 3 }
minikube service hello --url
```

### Invoke

```bash
SERVICE_URL="$(minikube service hello --url)"
curl -X POST -H 'Content-Type:application/json' "$SERVICE_URL/hello"
```

#### Swagger

```bash { terminalRows=3 }
SERVICE_URL="$(minikube service hello --url)"
echo "$SERVICE_URL/q/swagger-ui/"
```

### Check workflow instances

```bash { terminalRows=3 }
SERVICE_URL="$(minikube service hello --url)"
echo "$SERVICE_URL/q/dev/org.kie.kogito.kogito-quarkus-serverless-workflow-devui/workflowInstances"
```

## Create the budget increase workflow

```bash { terminalRows=4 }
kn workflow create --name serverless-travel-budget
```

### Verify

```bash { mimeType=text/x-json }
cat serverless-travel-budget/workflow.sw.json
```

### Edit and deploy the workflow

```bash { terminalRows=1 }
cp ./hack/workflow.sw.json serverless-travel-budget/workflow.sw.json
kn workflow deploy
code serverless-travel-budget/workflow.sw.json
```

### Invoke

```bash { mimeType=text/x-json }
SERVICE_URL="$(minikube service travelbudget --url)"
curl -X POST -H 'Content-Type:application/json' -d '{"currentBudget": 1000, "newBudget": 2800}' "$SERVICE_URL/travelbudget"
```

```bash { mimeType=text/x-json }
SERVICE_URL="$(minikube service travelbudget --url)"
curl -X POST -H 'Content-Type:application/json' -d '{"currentBudget": 1000, "newBudget": 1200}' "$SERVICE_URL/travelbudget"
```

### Get the service URL

```bash { terminalRows= 3}
minikube service travelbudget --url
```

#### Swagger

```bash { terminalRows=3 }
SERVICE_URL="$(minikube service travelbudget --url)"
echo "$SERVICE_URL/q/swagger-ui/"
```

### Check workflow instances

```bash { terminalRows=3 }
SERVICE_URL="$(minikube service travelbudget --url)"
echo "$SERVICE_URL/q/dev/org.kie.kogito.kogito-quarkus-serverless-workflow-devui/workflowInstances"
```

## Clean up

```bash
rm -rf demo/serverless-hello-world
rm -rf demo/serverless-travel-budget
```


```bash
minikube stop
```

or
```bash
minikube delete
```