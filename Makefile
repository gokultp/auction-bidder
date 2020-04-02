include ./build/auctionmanager/Makefile  ./build/usermanager/Makefile

build: build-users build-auctions

deploy-db:
	kubectl apply -f deploy/postgres/k8s.yaml

deploy: deploy-users deploy-auctions deploy-db

clean: clean-users clean-auctions

all: build deploy

