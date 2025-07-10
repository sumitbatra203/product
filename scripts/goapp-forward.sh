
kubectl port-forward -n gormapp $(kubectl get pod -n gormapp|grep go-app|awk '{print $1}') 8090:8080
