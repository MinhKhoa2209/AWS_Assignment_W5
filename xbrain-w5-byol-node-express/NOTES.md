# Deployment Notes

## Strategy

I chose the `serverless-http` adapter because it matches the README's minimum-change path:

- `app.js` remains framework-pure and Lambda-unaware.
- `server.js` remains the local HTTP runner for `npm start`.
- Lambda uses a small separate entrypoint, `lambda.js`, that adapts the existing Express app.
- The SAM template keeps the required HTTP API access method and only needs the handler value filled in.

## Smoke Test

- API URL: `https://wffmj4daqd.execute-api.us-west-2.amazonaws.com/`
- `GET /`: passed
- `GET /api/hello/Lan`: passed
- `POST /api/echo`: passed

## Cold Start

Measured from CloudWatch `REPORT` lines using:

```bash
sam logs --stack-name dmk-node-express --region us-west-2 -t
```

Observed invocations:

| Request | Duration | Billed Duration | Memory Size | Max Memory Used | Init Duration |
|---|---:|---:|---:|---:|---:|
| Cold start | `15.20 ms` | `289 ms` | `512 MB` | `93 MB` | `273.01 ms` |
| Warm start | `3.69 ms` | `4 ms` | `512 MB` | `93 MB` | N/A |

Representative cold start value for worksheet: `273.01 ms`.
