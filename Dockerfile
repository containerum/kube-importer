FROM golang:1.11-alpine as builder
RUN apk add --update make git
WORKDIR src/github.com/containerum/kube-importer
COPY . .
RUN VERSION=$(git describe --abbrev=0 --tags) make build-for-docker

FROM alpine:3.8

VOLUME ["/cfg"]

COPY --from=builder /tmp/kube-importer /

ENV KUBE_CONF="" \
    DEBUG=false \
    PORT=1666 \
    EXTLOG=true \
    CORS=true \
    EXCLUDED_NS="default,kube-system" \
    RESOURCE="http" \
    RESOURCE_URL="http://resource-service:1213" \
    PERMISSIONS="http" \
    PERMISSIONS_URL="http://permissions:4242" \
    VOLUMES="http" \
    VOLUMES_URL="http://volume-manager:4343"

EXPOSE 1666
CMD ["/kube-importer"]