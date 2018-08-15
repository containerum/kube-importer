FROM golang:1.10-alpine as builder
WORKDIR /go/src/github.com/containerum/kube-importer
COPY . .
WORKDIR cmd/kube-importer
RUN CGO_ENABLED=0 go build -v -ldflags="-w -s -extldflags '-static'" -tags="jsoniter" -o /bin/kube-importer

FROM alpine:3.7
COPY --from=builder /bin/kube-importer /
ENV CH_KUBE_IMPORTER_DEBUG="true" \
    CH_KUBE_IMPORTER_TEXTLOG="true"
VOLUME ["/cfg"]
EXPOSE 1212
CMD ["/kube-importer"]