FROM golang:1.23 AS go-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o build/simulog cmd/simulog/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=go-builder /app/build/simulog ./
ENTRYPOINT /app/simulog
