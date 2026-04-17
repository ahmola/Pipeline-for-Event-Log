# Build Stage
FROM --platform=$BUILDPLATFORM golang:1.25.6-alpine AS builder

ARG TARGETPLATFORM
ARG TARGETARCH
ARG TARGETOS
RUN echo "Building for $TARGETPLATFORM ($TARGETOS/$TARGETARCH)"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o main ./cmd/main

# Run Stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

ENV TZ=Asia/Seoul

RUN chmod +x ./main

CMD ["./main"]