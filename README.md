# dev-in-k8s

Provides some simple guidance for carrying out local development while your app runs within a Kubernetes cluster.

[Install kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
```
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.11.1/kind-linux-amd64
chmod +x ./kind
mv ./kind /usr/local/bin/kind
```

Install kubectl
```
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```

Create a kind cluster, using custom config with extra mounts so that the app source code is mounted into the kind node
```
kind create cluster --config ./kind.yaml
```

Build app image. Note that the app image [dlv](github.com/go-delve/delve/cmd/dlv) for debugging and [reflex](github.com/cespare/reflex) for watching source files and re-running the app.

```
docker build -f Dockerfile -t test-app:latest .
```

Confirm the app image works locally
```
docker run -it --rm -v $PWD:/work -p 8080:8080 test-app:latest 
```

Push the app image into kind, or a remote registry
```
kind load docker-image test-app:latest
```

Run the app pod, and port-forward (if you need)
```
kubectl apply -f ./deployment.yaml 
kubectl port-forward deployment/test-deployment 8080:8080
```

Make changes to your app source code locally and observe that reflex will reload it.

NOTE: for remote Kubernetes clusters you can remove the hostPath and the associated volume mount, and instead try to use [ksync](https://ksync.github.io/ksync/) for synchronising files.
