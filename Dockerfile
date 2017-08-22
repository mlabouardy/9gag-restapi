FROM golang:1.8.0

ADD . /go/src/github.com/mlabouardy/9gag-restapi

RUN go get github.com/gorilla/mux github.com/mlabouardy/9gag

EXPOSE 3000

WORKDIR /go/src/github.com/mlabouardy/9gag-restapi

CMD go run app.go
