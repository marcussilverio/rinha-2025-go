FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY main.go ./
RUN go build -o /main main.go

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=builder /main /main
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/main"]


