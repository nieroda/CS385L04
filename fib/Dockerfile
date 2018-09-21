FROM alpine as builder

RUN apk update && apk add go git gcc musl-dev
RUN mkdir ~/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:/go/bin:$PATH:/$GOROOT/bin
WORKDIR $GOPATH

COPY * /go/src/

RUN go get -ldflags "-linkmode external -extldflags -static" -a gofib
RUN go build -ldflags "-linkmode external -extldflags -static" -a gofib

FROM scratch

COPY --from=builder /go/bin/gofib /go/bin/gofib
CMD ["/go/bin/gofib"]
