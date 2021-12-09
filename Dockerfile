FROM golang:1.17 AS xcaddy-builder
RUN go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest

FROM golang:1.17 AS caddy-builder
COPY --from=xcaddy-builder /go/bin/xcaddy /bin/xcaddy
ADD ./ /go/src/cadmod
RUN xcaddy build --with github.com/dgsb/cadmod@latest=/go/src/cadmod --output /caddy

FROM ubuntu:20.04
COPY --from=caddy-builder /caddy /usr/local/bin/caddy
