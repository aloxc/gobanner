FROM golang
LABEL maintainer="leerohwa@gmail.com"
LABEL add="https://github.com/aloxc/gobanner"
LABEL version="1.0"
LABEL description="this is a golang banner,using golang file create a banner"

RUN go get github.com/aloxc/gobanner
WORKDIR /go/src/github.com/aloxc/gobanner/
RUN go build
RUN /go/src/github.com/aloxc/gobanner/gobanner
