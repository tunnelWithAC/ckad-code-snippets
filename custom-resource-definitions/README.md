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