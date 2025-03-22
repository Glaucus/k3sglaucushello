# Hello world deployment template

This repository aims to provide a easy to replicate baseline/template for other deployments on the Glaucus K3S
infrastructure

# How kubernetes files are handled?

Just place all the .yaml files on the `.k3s` folder all the deployment will execute the updates autmatically.

# How code changes are handled?

On every commit a new push to dockerhub is executed, and the pods are rotated automatically.

# Variable settings:

| Key              | Description                                               |
|------------------|-----------------------------------------------------------|
| KUBECONFIG       | K3S kubeconfig settings file encoded in base64            |
| DOCKER_PASSWORD  | DockerHub account password                                |
| DOCKER_USERNAME  | DockerHub account username                                |
| DOCKER_IMAGE_TAG | DockerHub image tag (username not included)               |
| K3S_NAMESPACE    | K3S namespace to execute updates on                       |
| K3S_DEPLOYMENT   | K3S deployment name                                       |

# Secrets management

In order to be super quick, right now the management is being done by GitHub secrets + plain environment variables

In order to use a secret, use a placeholder value in the kubernetes YAML files
and update the `.github` deployment to perform a sed replace.

Example secret:

| Key              | Description                                               |
| FANCY_SECRET     | Secret environment variable to be added at the deployment |

Kubernetes deployment env:

```yml
      containers:
        - name: k3sglaucushello
          image: nuriofernandez/k3sglaucushello:latest
          env:
            - name: SECRET
              value: "GITHUB_SECRET_FANCY_SECRET"
```

GitHub deploy sed:

```yml
      - name: Replace GITHUB_SECRET_ variables with actual GitHub Secrets
        run: |
          sed -i -E 's/GITHUB_SECRET_FANCY_SECRET/${{ secrets.FANCY_SECRET }}/g' .k3s/*.yaml
```

Usage and output:

```go
fmt.Fprintf(w, "K3S on Glaucus! With CI/CD! Fancy! "+os.Getenv("SECRET"))
```

![](https://i.imgur.com/Ko3aCdV.png)