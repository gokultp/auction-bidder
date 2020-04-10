build-worker:
	docker build -t worker -f build/worker/Dockerfile .

deploy-worker: 
	helm template  --name-template worker deploy/worker > deploy/worker/k8s.yaml
	kubectl apply  -f deploy/worker/k8s.yaml

clean-worker:
	kubectl delete -f deploy/worker/k8s.yaml