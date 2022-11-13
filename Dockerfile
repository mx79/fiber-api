FROM golang:1.19 AS builder
ENV GO111MODULE=on
WORKDIR /build
COPY . .
RUN go mod tidy
WORKDIR /build/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/app .

# Second stage for a smaller image
FROM alpine
COPY --from=builder /go/bin/app /go/bin/server/
WORKDIR /go/bin/server/
CMD ["./app"]