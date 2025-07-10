# SHELL := /bin/bash

# GOBIN := $(PWD)/bin
# PATH := $(GOBIN):$(PATH)

# #CLUSTER_NAME ?= kubejob-cluster
# KUBECONFIG ?= $(CURDIR)/.kube/config
# export KUBECONFIG
# export GOBIN

.PHONY: tidy
tidy: 
	go mod tidy

.PHONY: build-local
build-local: 
	CGO_ENABLED=0 GOOS=darwin go build -o build_out/go-app main.go
	docker build -t go-app .

.PHONY: build
build: 
	CGO_ENABLED=0 GOOS=linux go build -o build_out/go-app main.go
	docker build -t go-app .

.PHONY: run
run: 
	docker run go-app:latest
