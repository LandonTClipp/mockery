FROM golang:1.14-alpine as builder

COPY ./ /mockery
RUN cd /mockery && ls mockery && go install ./...

FROM golang:1.14-alpine

COPY --from=builder /go/bin/mockery /usr/local/bin

# Explicitly set a writable cache path when running --user=$(id -u):$(id -g)
# see: https://github.com/golang/go/issues/26280#issuecomment-445294378
ENV GOCACHE /tmp/.cache

ENTRYPOINT ["/usr/local/bin/mockery"]
