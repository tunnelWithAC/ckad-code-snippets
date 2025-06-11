As a good practice, take a backup of that apiserver manifest file before going to make any changes.

In case, if anything happens due to misconfiguration you can replace it with the backup file.

```
cp -v /etc/kubernetes/manifests/kube-apiserver.yaml /root/kube-apiserver.yaml.backup
```

Now, open up the kube-apiserver manifest file in the editor of your choice. It could be vim or nano.

```
vi /etc/kubernetes/manifests/kube-apiserver.yaml
```

Add the `--runtime-config` flag in the command field as follows :-

```
 - command:
    - kube-apiserver
    - --advertise-address=10.18.17.8
    - --allow-privileged=true
    - --authorization-mode=Node,RBAC
    - --client-ca-file=/etc/kubernetes/pki/ca.crt
    - --enable-admission-plugins=NodeRestriction
    - --enable-bootstrap-token-auth=true
    - --runtime-config=rbac.authorization.k8s.io/v1alpha1 --> This one 
```
    
After that kubelet will detect the new changes and will recreate the apiserver pod.

It may take some time.

```
kubectl get po -n kube-system
```
Check the status of the apiserver pod. It should be in running condition.
