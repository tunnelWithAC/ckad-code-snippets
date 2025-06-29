## Common Transformations

- commonLabel: adds a label to all Kubernetes resources
- namePrefix/Suffix: adds a common prefix/ suffix to all resource names
- Namespace: adds a common namespace to all resources
- commonAnnotations: adds an annotation to all resources


```
# kustomization.yaml

commonLabels:
    org: KodeKloud

namespace: lab

namePrefix: KodeKloud-

nameSuffix: -dev

commonAnnotations:
    branch: master
```