build-users:
	docker build -t usermanager -f build/usermanager/Dockerfile .


deploy-users: 
	helm template  --name-template usermanager deploy/usermanager > deploy/usermanager/k8s.yaml
	kubectl apply -f deploy/usermanager/k8s.yaml

clean-users:
	kubectl delete -f deploy/usermanager/k8s.yaml 