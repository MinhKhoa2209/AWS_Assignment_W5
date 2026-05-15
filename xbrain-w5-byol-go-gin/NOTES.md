# Deployment Notes

## Strategy

I chose `aws-lambda-go-api-proxy/gin` because it matches the README's minimum-change path for Gin on Lambda.

- `server/server.go` remains Lambda-unaware.
- `cmd/local/main.go` remains the local HTTP runner for `go run ./cmd/local`.
- Lambda uses a small root `main.go` entrypoint so SAM's `go1.x` build can produce the `bootstrap` binary.
- The Gin adapter handles API Gateway HTTP API v2 events and forwards them to the existing router.
