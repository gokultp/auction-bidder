build-scheduler:
	docker build -t scheduler -f build/scheduler/Dockerfile .

deploy-scheduler: 
	kubectl apply  -f deploy/scheduler/k8s.yaml

clean-scheduler:
	kubectl delete -f deploy/scheduler/k8s.yaml