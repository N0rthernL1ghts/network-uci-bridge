FROM --platform=${TARGETPLATFORM} golang:1.20-alpine AS builder

RUN set -eux \
    && apk add --no-cache make fish

COPY ["./cmd", "/project/network-uci-bridge/cmd"]
COPY ["./pkg", "/project/network-uci-bridge/pkg"]
COPY ["./go.mod", "./Makefile", "/project/network-uci-bridge/"]

WORKDIR "/project/network-uci-bridge"
RUN set -eux \
    && make build



FROM scratch AS rootfs

COPY --from=builder ["/project/network-uci-bridge/build/release/uci-bridge", "/usr/local/bin/"]



FROM --platform=${TARGETPLATFORM} alpine:3.18 AS final

COPY --from=rootfs ["/", "/"]

ENV UCI_TCP_HOST=""
ENV UCI_TCP_PORT="3333"

ENTRYPOINT ["/usr/local/bin/uci-bridge"]