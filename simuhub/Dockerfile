FROM golang:1.23 AS go-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY cmd/ cmd/
COPY internal/ internal/
RUN CGO_ENABLED=0 GOOS=linux go build -v -o build/simuhub cmd/simuhub/main.go

FROM node:23-alpine AS node-builder
WORKDIR /app
RUN npm install -g pnpm
COPY web/ ./
RUN pnpm install && pnpm build

FROM alpine:latest
WORKDIR /app
COPY --from=go-builder /app/build/simuhub ./
COPY --from=node-builder /app/build/ static/
EXPOSE 8080
ENTRYPOINT /app/simuhub
