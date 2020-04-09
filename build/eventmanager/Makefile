build-events:
	docker build -t eventmanager -f build/eventmanager/Dockerfile .

deploy-events: 
	helm template  --name-template eventmanager deploy/eventmanager > deploy/eventmanager/k8s.yaml
	kubectl apply  -f deploy/eventmanager/k8s.yaml

clean-events:
	kubectl delete -f deploy/eventmanager/k8s.yaml