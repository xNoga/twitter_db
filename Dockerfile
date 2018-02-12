FROM golang:jessie

RUN go get github.com/gin-gonic/gin
RUN go get gopkg.in/mgo.v2

WORKDIR /go/src/app

ADD src src

CMD [ "go", "run", "src/main.go" ]