ARG BUILD_GO_IMAGE=golang:1.18-alpine
ARG LIVE_GO_IMAGE=scratch

###
### Setup
###

### Setup Build
FROM ${BUILD_GO_IMAGE} AS setup-build

ARG GOARCH
ARG GOOS=linux

ENV GOARCH=${GOARCH}
ENV GOBIN=${GOPATH}/bin
ENV GOOS=${GOOS}

WORKDIR /usr/local/src

# Ease the debugging and development build by adding the go
# installed binaries to the PATH
ENV PATH=${PATH}:${GOBIN}

# Copy go.mod and go.sum files seperatly to allow caching
# of dependencies
COPY go.* ./
# Download dependencies
RUN go mod download

COPY . .

### Setup Live Image
FROM ${LIVE_GO_IMAGE} AS setup-live

COPY --from=setup-build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

###
### Build
###
FROM setup-build AS build

# Build the application binary
RUN CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${GOBIN}/terraform-config-inspect .

###
### Final Images
###

### Dev
# This stage is used for development and debugging purposes,
# but it can also be used for CI-related tasks, and others.
FROM setup-build AS dev

# Install delve debugger (the delve version shouldn't be a concern 🤷‍♂️)
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Explicitly not overriding the entrypoint of the source image
# ENTRYPOINT [ "" ]

### Live
FROM setup-live AS live

COPY --from=build /go/bin/terraform-config-inspect /

ENTRYPOINT ["/terraform-config-inspect"]
