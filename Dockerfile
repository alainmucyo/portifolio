FROM golang:1.14 as build
WORKDIR /go/src/github.com/alainmucyo/my-brand
COPY go.mod go.sum  ./
RUN GO111MODULE=on GOPROXY="https://proxy.golang.org" go mod download
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 go build -o /bin/my-brand .

FROM scratch
WORKDIR /
EXPOSE 5500

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /bin/my-brand /bin/my-brand

ENTRYPOINT ["/bin/my-brand"]