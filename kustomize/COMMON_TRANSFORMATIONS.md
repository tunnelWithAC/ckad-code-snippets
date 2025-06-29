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

## Image Transformer

Replace the image used in a deployment.

Name under `kustomization.yaml` specifies the name of the image and is linked to the Container `image` in a deployment rather than the Container `name`. 

```
images:
    - name: nginx
      newName: haproxy
```

You can also just change the tag instead of the image

```
images:
    - name: nginx
      newTag: 2.4
```

Or do both together

```
images:
    - name: nginx
      newName: haproxy
      newTag: 2.4
```