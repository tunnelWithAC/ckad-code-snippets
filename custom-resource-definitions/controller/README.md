# Custom Website Controller

This is a custom controller implementation for managing Website resources in Kubernetes. The controller watches for Website custom resources and manages their lifecycle.

## Prerequisites

- Go 1.16 or later
- Kubernetes cluster (v1.16 or later)
- kubectl configured to communicate with your cluster

## Project Structure

```
controller/
├── main.go         # Main controller implementation
├── go.mod          # Go module definition
└── README.md       # This file
```

## Implementation Details

The controller implements a custom resource type called `Website` with the following structure:

```go
type Website struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    Spec   WebsiteSpec   `json:"spec,omitempty"`
    Status WebsiteStatus `json:"status,omitempty"`
}

type WebsiteSpec struct {
    Domain   string `json:"domain"`
    Replicas int32  `json:"replicas"`
    Image    string `json:"image"`
    Port     int32  `json:"port"`
}

type WebsiteStatus struct {
    AvailableReplicas int32  `json:"availableReplicas"`
    Phase             string `json:"phase"`
}
```

## Required Dependencies

The controller requires the following Kubernetes client-go packages:

```go
k8s.io/apimachinery/pkg/apis/meta/v1
k8s.io/apimachinery/pkg/runtime
k8s.io/apimachinery/pkg/runtime/schema
k8s.io/client-go/dynamic
k8s.io/client-go/kubernetes
k8s.io/client-go/rest
k8s.io/client-go/tools/cache
k8s.io/client-go/tools/clientcmd
```

## Setup Instructions

1. Initialize the Go module:
   ```bash
   go mod init website-controller
   ```

2. Add required dependencies:
   ```bash
   go get k8s.io/client-go@latest
   go get k8s.io/apimachinery@latest
   ```

3. Build the controller:
   ```bash
   go build -o website-controller
   ```

4. Deploy the controller to your cluster:
   ```bash
   kubectl apply -f deployment.yaml
   ```

## Controller Features

- Watches for Website custom resources
- Handles create/update/delete events
- Supports both in-cluster and out-of-cluster configurations
- Graceful shutdown handling

## Implementation Notes

1. The controller uses a dynamic client to interact with the Kubernetes API server
2. It implements an informer pattern for efficient resource watching
3. The controller can run both inside and outside the cluster
4. Event handlers are implemented for Add, Update, and Delete operations

## TODO

- [ ] Implement reconciliation logic in event handlers
- [ ] Add proper error handling and retries
- [ ] Implement status updates
- [ ] Add metrics and monitoring
- [ ] Add proper logging and tracing
- [ ] Implement leader election for high availability

## References

- [Kubernetes Sample Controller](https://github.com/kubernetes/sample-controller)
- [Kubernetes Client-Go Documentation](https://github.com/kubernetes/client-go)
- [Custom Resource Definitions](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) 