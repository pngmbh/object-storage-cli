FROM quay.io/deis/go-dev:v1.17.2 as go-build

WORKDIR /go/src/github.com/deis/object-storage-cli
COPY . ./
RUN make bootstrap DEV_ENV_CMD=
RUN make build-all DEV_ENV_CMD= DIST_DIR=/opt

FROM scratch
COPY --from=go-build /opt /
