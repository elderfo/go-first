FROM golang


ENV GOPATH /opt/go:$GOPATH
ENV PATH /opt/go/bin:$PATH
ADD . /opt/go/src/local/myorg/go-first
WORKDIR /opt/go/src/local/myorg/go-first


RUN go get github.com/derekparker/delve/cmd/dlv
RUN go build
CMD ["./go-first"]