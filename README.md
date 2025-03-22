# Hello world deployment template

This repository aims to provide a easy to replicate baseline/template for other deployments on the Glaucus K3S infrastructure

# How kubernettes files are handled?

Just place all the .yaml files on the `.k3s` folder all the deployment will execute the updates autmatically.

# How code changes are handled?

On every commit a new push to dockerhub is executed, and the pods are rotated automatically.

# Variable settings:

| Key              | Description                                    |
|------------------|------------------------------------------------|
| KUBECONFIG       | K3S kubeconfig settings file encoded in base64 |
| DOCKER_PASSWORD  | DockerHub account password                     |
| DOCKER_USERNAME  | DockerHub account username                     |
| DOCKER_IMAGE_TAG | DockerHub image tag (username not included)    |
| K3S_NAMESPACE    | K3S namespace to execute updates on            |
| K3S_DEPLOYMENT   | K3S deployment name                            |
