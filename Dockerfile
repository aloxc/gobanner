FROM golang:1.12.4
LABEL maintainer="leerohwa@gmail.com"
LABEL add="https://github.com/aloxc/gobanner"
LABEL version="1.0"
LABEL description="this is a golang banner,using golang file create a banner"
RUN go get github.com/aloxc/gobanner
WORKDIR /go/src/github.com/aloxc/gobanner/
RUN go build && mv /go/src/github.com/aloxc/gobanner/gobanner /root \
    && rm -rf /go && rm -rf /usr/local/go \
    && apt-get remove -y g++ gcc libc6-dev make pkg-config \
    && rm -rf /var/lib/apt/lists/* \
    && /root/gobanner