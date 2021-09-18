
FROM golang:1.17 AS base

WORKDIR /go/src

COPY Makefile ./
COPY go.mod ./
COPY go.sum ./

RUN go mod download


COPY pkg pkg
COPY cmd cmd

RUN go build -o /go/bin/app cmd/app/*.go

FROM gcr.io/distroless/base-debian10

COPY --from=base /go/bin/app /

ENTRYPOINT ["/app"]