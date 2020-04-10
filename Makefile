include ./build/auctionmanager/Makefile  ./build/usermanager/Makefile ./build/scheduler/Makefile
include ./build/eventmanager/Makefile ./build/worker/Makefile  

build: build-users build-auctions build-scheduler build-events build-worker

deploy-db:
	kubectl apply -f deploy/postgres/k8s.yaml

deploy-queue:
	kubectl apply -f deploy/beanstalkd/k8s.yaml

deploy: deploy-users deploy-auctions deploy-events deploy-scheduler deploy-worker deploy-db deploy-queue

clean: clean-users clean-auctions clean-scheduler clean-events clean-worker

all: build deploy

