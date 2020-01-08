FROM golang:1.13

COPY . /build 

WORKDIR /build

RUN go build 


FROM alpine 

RUN apk add --no-cache ca-certificates

COPY --from=0 /build/go-socks5 /usr/bin/go-socks5

ENTRYPOINT [ "/usr/bin/go-socks5" ]