FROM golang:1.12 as builder
WORKDIR /auctions
COPY . /auctions/
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags "-w -s -X main.Version=${VERSION} -X main.MinVersion=$(git rev-parse HEAD) -X main.BuildTime=$(date +%FT%T%z)" ./cmd/eventmanager


FROM alpine:3.11 as deploy
WORKDIR /auctions

COPY --from=builder /auctions/eventmanager ./eventmanager
COPY --from=builder /auctions/rsa/id_rsa ./id_rsa
COPY --from=builder /auctions/rsa/id_rsa.pub ./id_rsa.pub

EXPOSE 80
CMD ["/auctions/eventmanager", "--port=80"]
