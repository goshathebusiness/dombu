FROM golang:1.21.1-alpine3.18 as build
WORKDIR /code
COPY . /code
RUN go build -buildvcs=false -o ./bin/user-management ./cmd/user-management/

#############
FROM alpine:3.18.3
WORKDIR /bin
COPY --from=build /code/bin/ ./

ENTRYPOINT ["./user-management"]