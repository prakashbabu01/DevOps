# Makefile for building and deploying the controller

IMG ?= your-docker-repo/pod-configmap-operator:latest

all: manager

# Build the manager binary
manager:
	go build -o bin/manager main.go

# Run against the Kubernetes cluster configured in ~/.kube/config
run: manifests generate
	go run ./main.go

# Deploy controller in the cluster
deploy: manifests generate
	kubectl apply -f config/default

# Build docker image
docker-build:
	docker build -t $(IMG) .

# Push docker image
docker-push:
	docker push $(IMG)
