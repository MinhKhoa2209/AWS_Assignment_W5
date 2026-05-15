# Deployment Notes

## Strategy

I chose `aws-lambda-go-api-proxy/gin` because it matches the README's minimum-change path for Gin on Lambda.

- `server/server.go` remains Lambda-unaware.
- `cmd/local/main.go` remains the local HTTP runner for `go run ./cmd/local`.
- Lambda uses a small root `main.go` entrypoint so SAM's `go1.x` build can produce the `bootstrap` binary.
- The Gin adapter handles API Gateway HTTP API v2 events and forwards them to the existing router.

## Smoke Test

- API URL: `https://exl1dzb081.execute-api.us-west-2.amazonaws.com/`
- `GET /`: passed
- `GET /api/hello/Lan`: passed
- `POST /api/echo`: passed

## Cold Start

Measured from CloudWatch `REPORT` lines using:

```bash
sam logs --stack-name dmk-go-gin --region us-west-2 -t
```

Observed cold starts:

| Request | Duration | Billed Duration | Memory Size | Max Memory Used | Init Duration |
|---|---:|---:|---:|---:|---:|
| Cold start | `1.96 ms` | `97 ms` | `256 MB` | `30 MB` | `94.47 ms` |
| Cold start | `2.55 ms` | `95 ms` | `256 MB` | `30 MB` | `92.15 ms` |
| Cold start | `2.10 ms` | `122 ms` | `256 MB` | `30 MB` | `119.77 ms` |

Representative cold start value for worksheet: `119.77 ms`.
