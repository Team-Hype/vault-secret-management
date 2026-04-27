# Vault Injector

## Overview

Vault Injector automatically injects secrets into Kubernetes pods using a sidecar container.

---

## How it works

1. Pod is created
2. Admission webhook modifies pod
3. Vault Agent sidecar is added
4. Agent authenticates with Vault
5. Secrets are written to files

---

## Configuration Steps

### Enable Kubernetes Auth

```
vault auth enable kubernetes
```

---

### Create Policy

```
path "secret/data/app/*" {
  capabilities = ["read"]
}
```

---

### Create Role

```
vault write auth/kubernetes/role/app \
  bound_service_account_names=default \
  bound_service_account_namespaces=default \
  policies=app-policy
```

---

## Pod Annotations

```
vault.hashicorp.com/agent-inject: "true"
vault.hashicorp.com/role: "app"
vault.hashicorp.com/agent-inject-secret-config.txt: "secret/data/app"
```

---

## Access Secrets

Inside pod:

```
/vault/secrets/config.txt
```
