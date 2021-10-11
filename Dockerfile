FROM golang:1.15-alpine as build

WORKDIR /go/src/github.com/Nikita99100/CarCRUDService

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
  -o bin/carservice \
  ./cmd

FROM alpine:3.12

COPY --from=build /go/src/github.com/Nikita99100/CarCRUDService/bin/carservice /usr/local/bin/carservice
COPY --from=build /go/src/github.com/Nikita99100/CarCRUDService/deployment/configs.toml /etc/carservice.toml

CMD ["/usr/local/bin/carservice", "-config", "/etc/carservice.toml"]