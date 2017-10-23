FROM golang:1.9.1-alpine3.6

EXPOSE 8080

WORKDIR /usr/src/app

RUN go get net/http \
	io \
	html/template

COPY . .

CMD ["go", "run", "app.go"]
