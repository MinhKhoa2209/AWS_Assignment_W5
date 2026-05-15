# Deployment Notes

## Strategy

I chose the `serverless-http` adapter because it matches the README's minimum-change path:

- `app.js` remains framework-pure and Lambda-unaware.
- `server.js` remains the local HTTP runner for `npm start`.
- Lambda uses a small separate entrypoint, `lambda.js`, that adapts the existing Express app.
- The SAM template keeps the required HTTP API access method and only needs the handler value filled in.
