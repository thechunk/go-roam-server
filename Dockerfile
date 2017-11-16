FROM golang

ADD . /go/src/github.com/thechunk/roam-server

RUN go get -u -v github.com/kardianos/govendor
RUN cd /go/src/github.com/thechunk/roam-server && govendor sync
RUN go install github.com/thechunk/roam-server

ENTRYPOINT /go/bin/roam-server

EXPOSE 8080
