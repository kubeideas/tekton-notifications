# Build the tekton notifications binary
FROM golang:1.18 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod


# Copy the go source
COPY main.go main.go
COPY email/ email/
COPY errors/ errors/
COPY slack/ slack/
COPY usage/ usage/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o notifications main.go

# Use distroless as minimal base image to package the notifications binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/notifications .
USER 65532:65532

ENTRYPOINT ["/notifications"]
