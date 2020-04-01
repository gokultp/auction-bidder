FROM golang:1.12 as builder
WORKDIR /auctions
COPY . /auctions/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -s -X main.Version=${VERSION} -X main.MinVersion=$(git rev-parse HEAD) -X main.BuildTime=$(date +%FT%T%z)" ./cmd/auctionmanager


FROM alpine:3.10 as deploy
WORKDIR /auctions

COPY --from=builder /auctions/auctionmanager ./auctionmanager

EXPOSE 80
CMD ["/auctions/auctionmanager", "--port=80"]
