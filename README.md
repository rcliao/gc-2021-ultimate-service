# GopherCon 2021 Workshop Exercise - Ultimate Service

This repository is created to capture my exercise for learning Ultimate Service (workshop in the GopherCon 2021)

## Practice Overview

This repository starts by setting up with:
1. A simple Go program (with console out) under `app` directory
2. Set up the Dockerfile under `zarf` so we can build out container of the main program
3. Set up the Kind and Kustomize file to run Kubernetes cluster locally
4. Utilize `Makefile` to group commands together

## Prerequisites
1. Install GoLang - https://go.dev
2. Install Kind - https://kind.sigs.k8s.io/docs/user/quick-start/
3. Install Kustomize - https://kustomize.io
4. Install Docker - https://www.docker.com
5. Install `kubectl` - https://kubectl.docs.kubernetes.io

## Get Started
1. Clone the repository
2. Run `make sales-api` to re-build the Docker image
3. Run `make kind-up` to set up the Kubernetes cluster
4. Run `make kind-load` to load the image in Kubernetes
5. Run `make kind-apply` to create pod/deployment
6. Run `make kind-status-sales` to check the container status - should be RUNNING
7. Run `make kind-logs-sales` to check the container logs - should be able to see last few lines of logs

