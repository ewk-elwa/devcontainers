## Tenant Capabilities Post-Setup
After RBAC configuration, tenants can:
1. Manage workloads: Create/delete pods, deployments, and jobs in their namespace.
1. Access namespaced resources: View and modify ConfigMaps, Secrets, and Services within their namespace.
1. Read logs: Access pod logs via kubectl logs.
1. Restricted cluster access: Cannot view or modify resources in other namespaces or cluster-scoped resources (e.g., Nodes, PersistentVolumes).


## Example Workflow for a Tenant
1. Deploy an application in their namespace:
```bash
kubectl apply -f deployment.yaml -n developers
```
1. View pods in their namespace:
```bash
kubectl get pods -n developers
```
1. Troubleshoot  in their namespace:
```bash
kubectl logs <pod-name> -n developers
```
