<img src="https://r2cdn.perplexity.ai/pplx-full-logo-primary-dark%402x.png" class="logo" width="120"/>

# Show example configuration of argocd agent to support two namespaces. Show now deployment from one namespace can not access the other

Below is an example configuration for a **single Argo CD agent** (instance) that supports deployments to two isolated namespaces—`team-one` and `team-two`—while ensuring that deployments from one namespace cannot access or modify resources in the other. This setup uses Argo CD **AppProjects** for logical isolation and Kubernetes RBAC for enforcement.

---

## **1. Label Namespaces for Argo CD Management**

Label each namespace so Argo CD is authorized to manage resources in them:

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: team-one
  labels:
    argocd.argoproj.io/managed-by: argocd
---
apiVersion: v1
kind: Namespace
metadata:
  name: team-two
  labels:
    argocd.argoproj.io/managed-by: argocd
```

*Here, `argocd` is the namespace where the Argo CD instance is running.*[^2]

---

## **2. Define Argo CD AppProjects for Each Team**

AppProjects restrict where applications can deploy resources. Example:

```yaml
apiVersion: argoproj.io/v1alpha1
kind: AppProject
metadata:
  name: team-one-project
  namespace: argocd
spec:
  destinations:
    - namespace: team-one
      server: https://kubernetes.default.svc
  sourceRepos:
    - '*'
---
apiVersion: argoproj.io/v1alpha1
kind: AppProject
metadata:
  name: team-two-project
  namespace: argocd
spec:
  destinations:
    - namespace: team-two
      server: https://kubernetes.default.svc
  sourceRepos:
    - '*'
```

*Each project only allows deployments to its designated namespace.*[^7][^8]

---

## **3. Example Argo CD Application for Each Namespace**

Team One application:

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: team-one-app
  namespace: argocd
spec:
  project: team-one-project
  source:
    repoURL: https://github.com/example/repo.git
    path: team-one-app
    targetRevision: HEAD
  destination:
    server: https://kubernetes.default.svc
    namespace: team-one
  syncPolicy:
    automated: {}
```

Team Two application (similar, but with `team-two-project` and `team-two` namespace).

---

## **4. Enforce Access Control with RBAC**

**Argo CD RBAC:**
Configure Argo CD RBAC so users or groups can only interact with their own AppProject and Applications. Example (in `argocd-rbac-cm` ConfigMap):

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: argocd-rbac-cm
  namespace: argocd
data:
  policy.csv: |
    p, role:team-one, applications, get, team-one-project/*, allow
    p, role:team-one, applications, sync, team-one-project/*, allow
    p, role:team-two, applications, get, team-two-project/*, allow
    p, role:team-two, applications, sync, team-two-project/*, allow
    # Deny cross-project access by omission
  policy.default: role:readonly
```

*Assign users to the appropriate roles via SSO or user management.*

---

## **5. How Isolation is Enforced**

- **AppProject destinations**: Each project restricts deployments to a single namespace. Even if a user tries to deploy to another namespace, Argo CD will reject the operation.
- **RBAC**: Users can only see and operate on their own applications and projects.
- **Kubernetes RBAC**: Further restricts service accounts so that Argo CD’s sync operations can only modify resources in the allowed namespaces.[^5][^6]

---

## **Result**

- **A deployment from `team-one` cannot access or modify resources in `team-two`**, and vice versa.
- Attempts to deploy to another team’s namespace will be blocked by Argo CD’s AppProject restrictions and RBAC policies.

---

**References:**

- [Argo CD: Applications in any namespace][^1]
- [Argo CD: Namespace Isolation][^6]
- [Argo CD: AppProject per namespace][^8]
- [Argo CD: ApplicationSet and multi-namespace support][^7]

<div style="text-align: center">⁂</div>

[^1]: https://argo-cd.readthedocs.io/en/latest/operator-manual/app-any-namespace/

[^2]: https://argocd-operator.readthedocs.io/en/latest/usage/deploy-to-different-namespaces/

[^3]: https://github.com/argoproj/argo-cd/issues/9508

[^4]: https://argo-cd.readthedocs.io/en/latest/user-guide/commands/argocd_admin_cluster_namespaces/

[^5]: https://github.com/argoproj/argo-cd/discussions/5193

[^6]: https://blog.andyserver.com/2020/12/argocd-namespace-isolation/

[^7]: https://argo-cd.readthedocs.io/en/latest/operator-manual/applicationset/Appset-Any-Namespace/

[^8]: https://github.com/argoproj/argo-cd/discussions/6255

