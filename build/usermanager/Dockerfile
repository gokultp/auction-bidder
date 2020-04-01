FROM golang:1.12 as builder
WORKDIR /users
COPY . /users/
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags "-w -s -X main.Version=${VERSION} -X main.MinVersion=$(git rev-parse HEAD) -X main.BuildTime=$(date +%FT%T%z)" ./cmd/usermanager


FROM alpine:3.10 as deploy
WORKDIR /users

COPY --from=builder /users/usermanager ./usermanager

EXPOSE 80
CMD ["/users/usermanager", "--port=80"]
