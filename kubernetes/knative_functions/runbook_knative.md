# Cluster Runbook

## Prerequisites

Install:

- `docker` or `podman`
- `kind` or `minikube`
- `kubectl` or `oc`
- `kn`

## Start cluster

### Kind

```bash
 kn quickstart kind --registry
```

Verify

```bash { terminalRows=3 }
kind get clusters
```

```bash
kubectl config use-context kind-knative
```

### Minikube

```bash { closeTerminalOnSuccess=true interactive=true }
kn quickstart minikube
```

Create the tunnel.

It is necessary because Knative uses `LoadBalancer` services to expose applications to the outside world.
The tunnel allows `minikube` to expose `LoadBalancer` services to your local computer, so that you can access them from your web browser.
_Knative Serving deployments_ and _Knative Eventing triggers_ are exposed as LoadBalancer services, so you need the tunnel to access them.

```bash { background=false }
minikube tunnel --profile knative
```

### Set the registry

```bash { promptEnv=false terminalRows=3 }
export FUNC_REGISTRY=quay.io/gciavarrini
```

## Write a Knative function

### Golang

#### Create

```bash { terminalRows=1 }
kn func create -l go hello-kn
```

```bash
tree hello-kn
```

#### Modify

```bash { closeTerminalOnSuccess=true interactive=false mimeType=text/x-go }
cat hello-kn/handle.go
```

```bash { mimeType=text/x-go }
cp ./handle_random.go ./hello-kn/handle.go
cat hello-kn/handle.go
```

### Build

```bash { name=func-build cwd=hello-kn }
kn func build
```

> NOTE: Make sure your registry is _public_

#### Deploy

```bash { closeTerminalOnSuccess=false cwd=hello-kn }
kn func deploy
```

#### Verify

```bash { terminalRows=4 }
kn func list
```

#### Invoke

```bash
kubectl get pods -w
```

```bash { cwd=hello-kn terminalRows=3 }
kn func invoke
```

### Quarkus

#### Create

```bash { terminalRows=25 }
 kn func create -l quarkus quarkus-hello-kn
 tree quarkus-hello-kn
```

```bash { mimeType=text/x-java terminalRows=25 }
cat quarkus-hello-kn/src/main/java/functions/Function.java
```

#### Deploy

```bash { cwd=quarkus-hello-kn terminalRows=5 }
kn func deploy
```

#### Verify

```bash { terminalRows=5 }
kn func list
```

#### Get info

```bash
kn func info quarkus-hello-kn
```

#### Invoke

```bash
kubectl get pods -w
```

```bash { cwd=quarkus-hello-kn terminalRows=2 }
kn func invoke
```

```bash { terminalRows=3 }
func invoke -h | grep -- "--data string"
```

```bash { cwd=quarkus-hello-kn terminalRows=6 }
echo "Invoke 1:"
kn func invoke --data "{\"message\": \"Today is LinuxDay\"}"
echo "\n\nInvoke 2"
kn func invoke --data "{\"message\": \"Tomorrow is Sunday\"}"
```

##### Configure scaling

```bash
cat > /tmp/config-autoscaler.yaml << EOF
apiVersion: v1
kind: ConfigMap
metadata:
  name: config-autoscaler
  namespace: knative-serving
data:
  min-scale: "2"
EOF
```

```bash
kubectl apply -f /tmp/config-scale.yml
```

##### cURL Alternative

```bash { closeTerminalOnSuccess=false terminalRows= }
QUARKUS_URL=$(kn func info quarkus-hello-kn | grep http | tr -d '[:blank:]')
curl ${QUARKUS_URL} \
  -H "Content-Type:application/json" \
  -d "{\"message\": \"Today is LinuxDay\"}"
echo "\n\nUse a different message\n"
curl ${QUARKUS_URL} \
  -H "Content-Type:application/json" \
  -d "{\"message\": \"Tomorrow is Sunday\"}"

```

# Clean up

```bash
kind delete cluster -n knative
rm -rf hello-kn
rm -rf quarkus-hello-kn
```