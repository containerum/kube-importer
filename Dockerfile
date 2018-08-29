FROM golang:1.11-alpine as builder
RUN apk add --update make git
WORKDIR src/github.com/containerum/kube-importer
COPY . .
RUN VERSION=$(git describe --abbrev=0 --tags) make build-for-docker

FROM alpine:3.8

VOLUME ["/cfg"]

COPY --from=builder /tmp/kube-importer /

ENV CH_KUBE_IMPORTER_KUBE_CONF="" \
    CH_KUBE_IMPORTER_DEBUG=false \
    CH_KUBE_IMPORTER_PORT=1666 \
    CH_KUBE_IMPORTER_TEXTLOG=true \
    CH_KUBE_IMPORTER_CORS=true \
    CH_KUBE_IMPORTER_EXCLUDED_NS="default,kube-system" \
    CH_KUBE_IMPORTER_RESOURCE="http" \
    CH_KUBE_IMPORTER_RESOURCE_URL="http://resource-service:1213" \
    CH_KUBE_IMPORTER_PERMISSIONS="http" \
    CH_KUBE_IMPORTER_PERMISSIONS_URL="http://permissions:4242" \
    CH_KUBE_IMPORTER_VOLUMES="http" \
    CH_KUBE_IMPORTER_VOLUMES_URL="http://volume-manager:4343"

EXPOSE 1666
CMD ["/kube-importer"]