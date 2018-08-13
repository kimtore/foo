FROM golang:alpine AS builder
COPY foo.go /foo.go
RUN apk --no-cache add git
RUN go get github.com/prometheus/client_golang/prometheus
RUN go build -o /foo /foo.go

FROM alpine
COPY --from=builder /foo /foo
COPY eicar-standard-antivirus-test-file.txt /eicar.txt
CMD ["/foo"]
