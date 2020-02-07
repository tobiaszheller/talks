ARG GO_IMAGE
ARG ALPINE_IMAGE
FROM golang:${GO_IMAGE}

ARG ARG_BUILD
ARG ARG_DATE
ARG ARG_BRANCH

COPY . /go/src/github.com/tobiaszheller/talks/black-friday/demo
RUN go install github.com/tobiaszheller/talks/black-friday/demo

FROM alpine:${ALPINE_IMAGE}
RUN apk add --no-cache ca-certificates

COPY --from=0 /go/bin/demo /demo
COPY . /go/src/github.com/tobiaszheller/talks/black-friday/demo
CMD ["/demo"]
