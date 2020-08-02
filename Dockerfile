FROM golang:1.13.10
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get github.com/mattn/go-sqlite3
RUN go get github.com/gorilla/mux
RUN go build -o main .
CMD ["/app/main"]