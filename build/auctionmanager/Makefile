build-auctions:
	docker build -t auctionmanager -f build/auctionmanager/Dockerfile .

deploy-auctions: 
	helm template  --name-template auctionmanager deploy/auctionmanager > deploy/auctionmanager/k8s.yaml
	kubectl apply  -f deploy/auctionmanager/k8s.yaml

clean-auctions:
	kubectl delete -f deploy/auctionmanager/k8s.yaml