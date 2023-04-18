FROM golang:alpine as builder

COPY . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api-gateway .

FROM alpine

WORKDIR /app
COPY --from=builder /app/api-gateway .
COPY config/*.yml ./config/
COPY migrations migrations
COPY wait-for .
CMD ["./api-gateway"]