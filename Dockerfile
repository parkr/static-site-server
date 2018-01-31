FROM golang:1.9.3

ADD . /go/src/github.com/parkr/static-site-server

RUN go install github.com/parkr/static-site-server

CMD ["/go/bin/static-site-server"]
