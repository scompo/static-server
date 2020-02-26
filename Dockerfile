FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -tags netgo -ldflags '-w' -o server server.go

FROM scratch
COPY --from=builder /build/server /app/
WORKDIR /app
CMD ["./server"]
