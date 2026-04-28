# Vault Basics

## Initialization

Vault must be initialized before use:

```
vault operator init
```

---

## Seal / Unseal

Vault starts sealed and must be unsealed to operate.

---

## Secrets Engines

* KV (key-value)
* PKI (certificates)
* Transit (encryption)

---

## Policies

Policies define access control rules.

---

## Authentication Methods

* Token
* Kubernetes
* AppRole
