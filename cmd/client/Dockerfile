FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /go/bin/app ./cmd/client

FROM scratch
COPY --from=builder /go/bin/app /go/bin/app
CMD ["/go/bin/app"]
