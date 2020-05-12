FROM golang:1.14

ARG importPath

WORKDIR $GOPATH/src/${importPath}

COPY . .
