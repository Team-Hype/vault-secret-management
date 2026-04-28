# Setup Guide

## Prerequisites

* kubectl
* helm
* Terraform
* Yandex Cloud CLI

---

## Kubernetes Access

```
yc config set service-account-key key.json
yc managed-kubernetes cluster get-credentials --id <cluster-id> --external
```

---

## Deploy Vault

```
helm upgrade --install vault hashicorp/vault \
  -n vault --create-namespace \
  -f vault/values.yaml
```

---

## Initialize Vault

```
kubectl exec -it vault-0 -n vault -- vault operator init
```

---

## Unseal Vault

Repeat 3 times:

```
vault operator unseal
```

---

## Enable KV Secrets Engine

```
vault secrets enable -path=secret kv-v2
```

---

## Store Secret

```
vault kv put secret/app/config username=admin password=123
```
