FROM golang:1.19 AS builder

WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN go build -o manager main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/manager .
USER 65532:65532

ENTRYPOINT ["/manager"]
