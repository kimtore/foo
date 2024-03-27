FROM golang:alpine AS builder
COPY . /src
RUN apk --no-cache add git
WORKDIR /src
RUN go build -o /foo /src/foo.go

FROM alpine
COPY --from=builder /foo /foo
CMD ["/foo"]
