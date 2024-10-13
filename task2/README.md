### 创建一个新的自定义 CRD（Group：aiops.geektime.com, Version: v1alpha1, Kind: AIOps），并使用 dynamicClient 获取该资源

```
kubectl apply -f crd.yaml
customresourcedefinition.apiextensions.k8s.io/aiops.aiops.geektime.com created
```

```
kubectl apply -f aiops.yaml
aiops.aiops.geektime.com/aiops-instance created
```

```
kubectl get aiops
NAME             AGE
aiops-instance   20s
```

```
go mod init task2
go: creating new go.mod: module task2
```

create and write main.go

```
go mod tidy
go: finding module for package k8s.io/client-go/util/homedir
go: finding module for package k8s.io/client-go/dynamic
go: finding module for package k8s.io/client-go/restmapper
go: finding module for package k8s.io/client-go/kubernetes
go: finding module for package k8s.io/client-go/tools/clientcmd
go: finding module for package k8s.io/apimachinery/pkg/apis/meta/v1
go: finding module for package k8s.io/apimachinery/pkg/runtime/schema
go: found k8s.io/apimachinery/pkg/apis/meta/v1 in k8s.io/apimachinery v0.31.1
go: found k8s.io/apimachinery/pkg/runtime/schema in k8s.io/apimachinery v0.31.1
go: found k8s.io/client-go/dynamic in k8s.io/client-go v0.31.1
go: found k8s.io/client-go/kubernetes in k8s.io/client-go v0.31.1
go: found k8s.io/client-go/restmapper in k8s.io/client-go v0.31.1
go: found k8s.io/client-go/tools/clientcmd in k8s.io/client-go v0.31.1
go: found k8s.io/client-go/util/homedir in k8s.io/client-go v0.31.1
```

```
go run main.go
Namespace: default, Name: aiops-instance
```