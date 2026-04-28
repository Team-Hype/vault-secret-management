# Vault + Kubernetes + Injector

## Overview

This project demonstrates secure secrets management using HashiCorp Vault integrated with Kubernetes.

The system allows applications to dynamically retrieve secrets without storing them in source code, Kubernetes manifests, or CI pipelines.

---

## Architecture

* Kubernetes cluster (Yandex Cloud)
* HashiCorp Vault deployed via Helm
* Vault Agent Injector for secret delivery
* GitHub Actions for CI/CD

Secrets flow:
Pod → Vault Agent → Vault → Secret file inside container

---

## Features

* Vault deployment using Helm
* Vault initialization and unsealing (lock/unlock mechanism)
* Secure storage of:

  * passwords
  * API tokens
  * private keys
  * certificates
* Kubernetes authentication
* Automatic secret injection into pods
* CI/CD automation with GitHub Actions

---

## Vault Lock / Unlock

Vault starts in a sealed state.

To initialize:

```
vault operator init
```

To unseal:

```
vault operator unseal
```

To seal again:

```
vault operator seal
```

---

## Setup

### 1. Provision infrastructure

```
cd terraform
terraform apply
```

### 2. Deploy Vault

Run GitHub Actions workflow:

* Deploy Vault

### 3. Initialize Vault

```
kubectl exec -it vault-0 -n vault -- vault operator init
```

Save:

* unseal keys
* root token

### 4. Unseal Vault

```
vault operator unseal
```

---

## Secrets Example

```
vault kv put secret/app username=admin password=secret
```

---

## Injector Demo

Deploy application:

```
kubectl apply -f k8s/app/deployment.yaml
```

Check injected secrets:

```
kubectl exec -it <pod> -- cat /vault/secrets/config.txt
```

---

## CI/CD

Workflows:

* Deploy Vault
* Deploy application

Both are triggered manually via GitHub Actions.

---

## Tech Stack

* Kubernetes
* HashiCorp Vault
* Helm
* Terraform
* GitHub Actions

---

## Demo Scenario

1. Deploy Vault
2. Initialize and unseal
3. Store secrets
4. Deploy application
5. Verify secret injection

---

## Documentation

See `docs/` folder for detailed setup and configuration.
