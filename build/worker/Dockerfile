FROM golang:1.12 as builder
WORKDIR /auctions
COPY . /auctions/
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags "-w -s -X main.Version=${VERSION} -X main.MinVersion=$(git rev-parse HEAD) -X main.BuildTime=$(date +%FT%T%z)" ./cmd/worker


FROM alpine:3.11 as deploy
WORKDIR /auctions

COPY --from=builder /auctions/worker ./worker


EXPOSE 80
CMD ["/auctions/worker", "--port=80"]
