FROM golang:1.12 as builder
WORKDIR /users
COPY . /users/
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags "-w -s -X main.Version=${VERSION} -X main.MinVersion=$(git rev-parse HEAD) -X main.BuildTime=$(date +%FT%T%z)" ./cmd/usermanager


FROM alpine:3.11 as deploy
WORKDIR /users

COPY --from=builder /users/usermanager ./usermanager
COPY --from=builder /users/rsa/id_rsa ./id_rsa
COPY --from=builder /users/rsa/id_rsa.pub ./id_rsa.pub



EXPOSE 80
CMD ["/users/usermanager", "--port=80"]
