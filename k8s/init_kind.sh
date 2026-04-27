NAMESPACE="event-system"
# 경로는 루트 디렉토리 pipeline
# 기본 클러스터 생성
kind create cluster --name $NAMESPACE

# 로컬 이미지를 kind 노드에 로드
kind load docker-image event-generator:latest --name $NAMESPACE

# 메니페스트 적용
kubectl apply -f ./k8s/manifests
# ArgoCD Manifest
kubectl apply -n $NAMESPACE -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# 파드 상태 확인 대기
while [ -z "$(kubectl get pods -n $NAMESPACE --no-headers 2>/dev/null)" ]; do
  echo "Waiting for pods to be initialized..."
  sleep 2
done
echo "Intialized Pods!"

kubectl wait --for=condition=Ready pod --all -n $NAMESPACE --timeout=-1s
echo "All pods are running!"

# ArgoCD 초기 비밀번호
kubectl get secret -n $NAMESPACE argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d

# 사용자에게 실행 여부 확인
read -p "Do you want to forward Grafana? [y/n]: " answer

# 입력값이 y 또는 Y인 경우에만 실행
if [[ "$answer" == "y" || "$answer" == "Y" ]]; then
    echo "Starting Grafana port-forwarding..."
    echo "If you want to stop port-forwarding, press Ctrl+C"
    kubectl port-forward svc/grafana 3000:3000 -n $NAMESPACE
else
    echo "Port-forwarding cancelled."
fi