# **Secure Secret Management Using HashiCorp Vault with Kubernetes Integration**


# **1. Introduction**

Modern cloud-native applications require secure handling of sensitive data such as API tokens, database credentials, and TLS certificates. Hardcoding or storing secrets in plain text introduces **serious** security risks.

This project addresses these issues by implementing:
* A **HashiCorp Vault** instance for secret management
* A **Kubernetes cluster** for orchestration
* A **Go-based application** that consumes secrets securely
* CI/CD automation via GitHub Actions
### **Objectives**
* Demonstrate Vault deployment and initialization (lock/unlock lifecycle)
* Store multiple types of secrets (passwords, tokens, certificates)
* Integrate Vault with Kubernetes using secret injection
* Provide a working real-world use case

---

# **2. Methods**

## **2.1 System Architecture**

The system consists of the following components:

```
+-------------------+
|   GitHub Actions  |
| (CI/CD Pipeline)  |
+--------+----------+
         |
         v
+-------------------+
|   Kubernetes      |
|   Cluster         |
|                   |
|  +-------------+  |
|  |   Vault     |<------------------+
|  | (Secrets)   |                   |
|  +------+------+                   |
|         |                          |
|         v                          |
|  +-------------+                   |
|  | Application |-------------------+
|  | (Go App)    |  (Injected secrets)
|  +-------------+
|
+-------------------+
```


---

## **2.2 Vault Deployment**

Vault is deployed using Helm with configuration defined in:
* `manifests/vault/values.yaml`

**Key features:**
* Dev/standalone mode setup
* Auto-unseal (if configured)
* Kubernetes authentication backend is enabled in Vault to allow authentication using Kubernetes ServiceAccounts.

---

## **2.3 Vault Initialization and Lifecycle**

Vault follows a strict security lifecycle:
### **Initialization**
*Vault Generates:*
  * Unseal keys (5)
  * Root token
### **Sealing/Unsealing**
* Vault starts in a **sealed state**
* Requires **unseal keys** to decrypt storage

```
Vault State Flow:

[ Sealed ] ---> (Unseal Keys) ---> [ Unsealed ]
     ^                                 |
     |----------(Seal Command)---------|
```

---

## **2.4 Secret Storage**

### **Examples implemented:**
* Database credentials (PostgreSQL)
* Tokens (CI/CD or service tokens)
* TLS certificates (via cert-manager)
* Application configuration

Secrets are stored using KV engine:
```
secret/data/postgres/dev
secret/data/postgres/prod
```
***Secrets are logically separated into development and production environments to enforce isolation.***

---

## **2.5 Kubernetes Integration**
Access to secrets is controlled using Vault roles bound to Kubernetes ServiceAccounts.
Each environment (dev and prod) has a dedicated ServiceAccount and Vault role.
### **1. Service Accounts**
Defined in:
* `manifests/database/serviceaccount-*.yaml`
### **2. Vault Agent Injector**
Described in:
* `docs/injector.md`
This enables automatic secret injection into pods


## **2.5.1 Access Control Model**

**Two isolated roles are defined:**
- postgres-dev → access to secret/data/postgres/dev
- postgres-prod → access to secret/data/postgres/prod

**Each role is strictly bound to a Kubernetes ServiceAccount:**
- postgres-dev-sa
- postgres-prod-sa

Vault enforces least-privilege access between environments.
**Secret injection and access control validation using HashiCorp Vault: dev and prod environments are strictly isolated**
## Secret Flow Diagram

```
Pod → Kubernetes ServiceAccount JWT → Vault Kubernetes Auth → Vault Token → Vault Agent Injector → Secret file in container
```

---

## **2.6 Application Implementation**

A Go application (`app/`) demonstrates secret consumption.
### **Key Components:**
* `config.go` → Loads configuration from environment
* `postgres.go` → Connects to DB using Vault-provided credentials
* `main.go` → Entry point displaying app info

Secrets are **not hardcoded**, but injected dynamically at runtime.

---

## **2.7 CI/CD Integration**

GitHub Actions workflows:
* `.github/workflows/deploy-vault.yaml`
* `.github/workflows/deploy-app.yaml`

Pipeline responsibilities:
* Deploy Vault
* Deploy application
* Apply Kubernetes manifests

---

# **3. Results**

## **3.1 Successful Vault Deployment**
* Vault initialized and unsealed
* Secrets stored securely
* Access control validation
## **3.2 Secret Injection Works**
*Application receives:*
  * DB credentials
  * Config values
*No secrets stored in code or repo*
## **3.3 Application Connectivity**
* Go service successfully connects to PostgreSQL
* Uses Vault-managed credentials
## **3.4 Automation Achieved**
* CI/CD pipeline deploys full system
* Reproducible infrastructure

---

## **3.5 Example Workflow**

```
1. Vault initialized
2. Secrets stored (DB creds, tokens)
3. Kubernetes pod starts
4. Vault Agent injects secrets
5. Application reads secrets
6. App connects to database
```

---

# **4. Discussion**

## **4.1 Advantages**
* **Security**: No plaintext secrets in code
* **Centralization**: All secrets in one system
* **Dynamic secrets**: Can rotate credentials
* **Scalability**: Works across environments

## **4.2 Limitations**
* Initial setup complexity
* Requires careful key management
* Vault availability is critical

## **4.3 Possible Improvements**
* Enable auto-unseal using cloud KMS
* Add secret rotation policies
* Integrate with:
	   GitLab CI (instead of GitHub Actions)
	   Ansible for provisioning

---

# **5. Conclusion**

This project successfully demonstrates how HashiCorp Vault can be used to securely manage secrets in a Kubernetes environment.

**Key achievements:**
* Vault deployment and lifecycle management
* Secure storage of multiple secret types
* Seamless integration with Kubernetes
* Real application consuming secrets dynamically

The solution reflects best practices for modern DevOps and software security.

---

# **6. References**
* HashiCorp Vault Documentation
* Kubernetes Documentation
* Vault Agent Injector Guide

---

# **Appendix**

## **A. Key Files Overview**

| File              | Purpose               |
| ----------------- | --------------------- |
| `values.yaml`     | Vault configuration   |
| `deployment.yaml` | Kubernetes deployment |
| `config.go`       | App config loader     |
| `postgres.go`     | DB connection         |
| GitHub workflows  | CI/CD automation      |

---

## **B. Diagram: Secret Injection**

```
+----------------------+
| Kubernetes Pod       |
|                      |
|  +----------------+  |
|  | Vault Agent    |  |
|  +--------+-------+  |
|           |          |
|           v          |
|  +----------------+  |
|  | Application    |  |
|  | (reads secret) |  |
|  +----------------+  |
|                      |
+----------+-----------+
           |
           v
     +-----------+
     |  Vault    |
     +-----------+
```

---

## **C. Diagram: CI/CD Flow**

```
Developer Push
      |
      v
GitHub Actions
      |
      v
Deploy Vault ---> Deploy App ---> Apply Manifests
      |
      v
Kubernetes Cluster
```

---


