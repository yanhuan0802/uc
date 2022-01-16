# 编写Kubernetes部署脚本将httpserver部署到kubernetes集群
### 优雅启动和优雅终止
* https://kubernetes.io/zh/docs/concepts/containers/container-lifecycle-hooks/
#### 容器终止流程
1. Pod 被删除，状态置为 Terminating。
2. kube-proxy 更新转发规则，将 Pod 从 service 的 endpoint 列表中摘除掉，新的流量不再转发到该 Pod。
3. 如果 Pod 配置了 preStop Hook ，将会执行。
4. kubelet 对 Pod 中各个 container 发送 `SIGTERM` 信号以通知容器进程开始优雅停止。
5. 等待容器进程完全停止，如果在 terminationGracePeriodSeconds 内 (默认 30s) 还未完全停止，就发送 `SIGKILL` 信号强制杀死进程。
6. 所有容器进程终止，清理 Pod 资源。
### 资源需求和QoS保证 
* https://kubernetes.io/zh/docs/tasks/configure-pod-container/quality-service-pod/
### 探活 
* https://kubernetes.io/zh/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/
### 日志等级、配置和代码分离
* https://kubernetes.io/zh/docs/concepts/configuration/configmap/
## Part2
### service
* https://kubernetes.io/zh/docs/concepts/services-networking/service/
### ingress
* https://kubernetes.io/zh/docs/concepts/services-networking/ingress/
* Deploy NGINX Ingress controller
```shell
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install ingress-nginx ingress-nginx/ingress-nginx
```
### 通过证书保证httpServer的通讯安全
* Deploy cert-manager
```shell
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.6.1/cert-manager.yaml
kubectl apply -f ./issuer.yaml
```

