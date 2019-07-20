FROM golang

ADD . /go/src/calc-service

RUN go get github.com/gorilla/mux
RUN go install calc-service

ENTRYPOINT /go/bin/calc-service

EXPOSE 8080