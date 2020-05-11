FROM golang:1.14

ARG importPath

WORKDIR /go/src/${importPath}
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8000

CMD ["favorite-tree"]
