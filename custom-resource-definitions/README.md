# Kubernetes Custom Resource Definitions (CRDs)

Custom Resource Definitions (CRDs) are extensions of the Kubernetes API that allow you to create custom resources. This enables you to define your own API objects and controllers to manage them.

## What are CRDs?

CRDs allow you to:
- Define custom resources that behave like native Kubernetes resources
- Create custom controllers to manage these resources
- Extend Kubernetes functionality without modifying the core codebase

## Example Implementation

This repository contains an example CRD implementation for a `Website` resource that manages web applications. The implementation includes:

1. CRD definition (YAML)
2. Go implementation of the custom controller
3. Example usage

## Structure

```
custom-resource-definitions/
├── README.md
├── crd.yaml
└── controller/
    └── main.go
```

## Getting Started

1. Apply the CRD:
```bash
kubectl apply -f crd.yaml
```

2. Build and run the controller:
```bash
go build -o controller
./controller
```

3. Create a custom resource:
```bash
kubectl apply -f example-website.yaml
```

## Best Practices

1. Use meaningful names for your CRDs
2. Implement proper validation
3. Follow Kubernetes API conventions
4. Include proper documentation
5. Implement proper error handling
6. Use finalizers for cleanup

## Kubernetes Operator Framework

The Kubernetes Operator Framework is a toolkit for building and managing Kubernetes native applications, called Operators. Operators are software extensions to Kubernetes that make use of custom resources to manage applications and their components.

### Key Components

1. **Operator SDK**
   - Provides tools to build, test, and package Operators
   - Supports multiple programming languages (Go, Ansible, Helm)
   - Includes scaffolding tools for quick Operator development

2. **Operator Lifecycle Manager (OLM)**
   - Manages the lifecycle of Operators in a cluster
   - Handles installation, upgrades, and dependency management
   - Provides a catalog of available Operators

3. **Operator Metering**
   - Tracks Operator usage and resource consumption
   - Helps with billing and resource planning

### Implementation Approaches

1. **Go-based Operators**
   ```go
   // Example using operator-sdk
   type WebsiteReconciler struct {
       client.Client
       Scheme *runtime.Scheme
   }

   func (r *WebsiteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
       var website v1alpha1.Website
       if err := r.Get(ctx, req.NamespacedName, &website); err != nil {
           return ctrl.Result{}, client.IgnoreNotFound(err)
       }
       // Reconciliation logic here
       return ctrl.Result{}, nil
   }
   ```

2. **Ansible-based Operators**
   ```yaml
   # Example playbook
   - name: Reconcile Website
     k8s:
       state: present
       definition:
         apiVersion: apps/v1
         kind: Deployment
         metadata:
           name: "{{ website.metadata.name }}"
         spec:
           replicas: "{{ website.spec.replicas }}"
   ```

3. **Helm-based Operators**
   ```yaml
   # Example values.yaml
   website:
     domain: "{{ .Values.domain }}"
     replicas: {{ .Values.replicas }}
     image: "{{ .Values.image }}"
   ```

### Best Practices

1. **Resource Management**
   - Use finalizers for proper cleanup
   - Implement proper status updates
   - Handle resource versioning

2. **Error Handling**
   - Implement retry mechanisms
   - Use proper error types
   - Log errors appropriately

3. **Testing**
   - Unit tests for business logic
   - Integration tests with the Kubernetes API
   - End-to-end tests for complete workflows

4. **Security**
   - Implement RBAC properly
   - Use service accounts
   - Follow security best practices

### Example Implementation Steps

1. **Initialize Operator Project**
   ```bash
   operator-sdk init --domain example.com --repo github.com/example/website-operator
   ```

2. **Create API and Controller**
   ```bash
   operator-sdk create api --group website --version v1alpha1 --kind Website
   ```

3. **Implement Reconciliation Logic**
   ```go
   func (r *WebsiteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
       // 1. Fetch the Website instance
       // 2. Check if the Website is being deleted
       // 3. Create/update required resources
       // 4. Update status
       // 5. Return result
   }
   ```

4. **Build and Deploy**
   ```bash
   make docker-build IMG=example/website-operator:latest
   make deploy IMG=example/website-operator:latest
   ```

### Additional Resources

- [Operator SDK Documentation](https://sdk.operatorframework.io/)
- [Operator Lifecycle Manager Documentation](https://olm.operatorframework.io/)
- [Kubernetes Operator Pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)
- [Operator Best Practices](https://sdk.operatorframework.io/docs/best-practices/)

## Usage

1. Apply the CRD:
   ```bash
   kubectl apply -f crd.yaml
   ```

2. Create a Website resource:
   ```bash
   kubectl apply -f example-website.yaml
   ```

3. Deploy the controller:
   ```bash
   kubectl apply -f controller/deployment.yaml
   ```

## Development

To modify or extend the controller:

1. Update the CRD definition in `crd.yaml`
2. Modify the controller code in the `controller/` directory
3. Rebuild and redeploy the controller

## Testing

Test the CRD and controller:

1. Apply the CRD
2. Create a Website resource
3. Verify the controller creates the necessary resources
4. Check the status of the Website resource 