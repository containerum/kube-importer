FROM golang:1.10-alpine as builder
RUN apk add --update make git
WORKDIR src/github.com/containerum/kube-importer
COPY . .
RUN VERSION=$(git describe --abbrev=0 --tags) make build-for-docker

FROM alpine:3.7

VOLUME ["/cfg"]

COPY --from=builder /tmp/kube-importer /

ENV CH_KUBE_IMPORTER_DEBUG="true" \
    CH_KUBE_IMPORTER_TEXTLOG="true"

EXPOSE 1212
CMD ["/kube-importer"]