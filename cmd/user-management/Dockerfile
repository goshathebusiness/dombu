FROM golang:1.21.1 as builder

ARG ARCH=amd64
ARG GO111MODULE=on

WORKDIR /code

COPY go.mod go.sum vendor /code/
COPY . /code
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build \
		-mod=vendor \
		-o $GOPATH/bin/user-management \
		-gcflags "all=-trimpath=$GOPATH" \
		-a ./cmd/user-management

# Run the container
FROM alpine:3.15

COPY --from=builder /go/bin/user-management /bin/user-management
ENTRYPOINT ["/bin/user-management"]