### 使用 Informer + RateLimitingQueue 监听 Pod 事件

```
go mod init task1
go: creating new go.mod: module task1
```

create and write main.go

```
go mod tidy
go: finding module for package k8s.io/client-go/util/workqueue
go: finding module for package k8s.io/api/core/v1
go: finding module for package k8s.io/client-go/informers
go: finding module for package k8s.io/client-go/kubernetes
go: finding module for package k8s.io/client-go/tools/cache
go: finding module for package k8s.io/client-go/tools/clientcmd
go: found k8s.io/api/core/v1 in k8s.io/api v0.31.1
go: found k8s.io/client-go/informers in k8s.io/client-go v0.31.1
go: found k8s.io/client-go/kubernetes in k8s.io/client-go v0.31.1
go: found k8s.io/client-go/tools/cache in k8s.io/client-go v0.31.1
go: found k8s.io/client-go/tools/clientcmd in k8s.io/client-go v0.31.1
go: found k8s.io/client-go/util/workqueue in k8s.io/client-go v0.31.1
```

```
go run main.go
Sync/Add/Update pod argocd-application-controller-0, status: Running
Sync/Add/Update pod argocd-applicationset-controller-5bc7f4ff55-vtnvn, status: Running
Sync/Add/Update pod argocd-dex-server-5b7c4f9d4f-h8qwk, status: Running
Sync/Add/Update pod argocd-notifications-controller-5cbc66fd96-w97rp, status: Running
Sync/Add/Update pod argocd-redis-7cdbbb8576-nf4zx, status: Running
Sync/Add/Update pod argocd-repo-server-f7b9c9859-gss5p, status: Running
Sync/Add/Update pod argocd-server-f9cf5db6c-qgjp2, status: Running
Sync/Add/Update pod crossplane-79775db9c8-7mgjc, status: Running
Sync/Add/Update pod crossplane-rbac-manager-66998d5c5d-6pgr4, status: Running
Sync/Add/Update pod provider-tencentcloud-8bf0507521fb-5fd5bc47dd-lkvrz, status: Running
Sync/Add/Update pod coredns-76f75df574-226np, status: Running
Sync/Add/Update pod coredns-76f75df574-j8pvm, status: Running
Sync/Add/Update pod etcd-docker-desktop, status: Running
Sync/Add/Update pod kube-apiserver-docker-desktop, status: Running
Sync/Add/Update pod kube-controller-manager-docker-desktop, status: Running
Sync/Add/Update pod kube-proxy-6fmjq, status: Running
Sync/Add/Update pod kube-scheduler-docker-desktop, status: Running
Sync/Add/Update pod storage-provisioner, status: Running
Sync/Add/Update pod vpnkit-controller, status: Running
Sync/Add/Update pod nginx-task1-5477b4ff8c-w446l, status: Pending
Sync/Add/Update pod nginx-task1-5477b4ff8c-w446l, status: Pending
Sync/Add/Update pod nginx-task1-5477b4ff8c-w446l, status: Pending
Sync/Add/Update pod nginx-task1-5477b4ff8c-w446l, status: Running
Sync/Add/Update pod nginx-task1-5477b4ff8c-w446l, status: Running
Sync/Add/Update pod nginx-task1-5477b4ff8c-w446l, status: Succeeded
Sync/Add/Update pod nginx-task1-5477b4ff8c-w446l, status: Succeeded
Sync/Add/Update pod nginx-task1-5477b4ff8c-w446l, status: Succeeded
Pod default/nginx-task1-5477b4ff8c-w446l does not exist anymore
```
