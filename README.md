# Deploy Nginx in K8s using this CLI tool

## Build
```shell
$ go build
```

## Usage
### Scale the deployment - 
```shell
$ deploynginx scale --replicas=2
```

### Update the version
```shell
$ deploynginx upgrade --version nginx:1.25.0
```

`Note:` This does not make any new deployment (as per requirement) please make sure you have
the deployment already created using `nginx.yaml` file.

