FROM golang:1.9.1-alpine3.6

EXPOSE 8080

RUN apk --update add git openssh && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

WORKDIR /usr/src/app

RUN go get net/http \
	io \
	html/template \
	github.com/julienschmidt/httprouter \
	gopkg.in/mgo.v2/bson

COPY . .

CMD ["go", "run", "app.go"]
