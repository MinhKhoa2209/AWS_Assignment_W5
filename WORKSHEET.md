# Worksheet

## Node.js Express

- Strategy: `serverless-http`
- Lambda function name: `dmk-node-express`
- Stack name: `dmk-node-express`
- HTTP access method: API Gateway HTTP API via SAM `HttpApi`
- Smoke test endpoint: `https://wffmj4daqd.execute-api.us-west-2.amazonaws.com/`
- Smoke test result: passed for `/`, `/api/hello/Lan`, and `/api/echo`
- Observed cold start: `273.01 ms`
- Warm start duration: `3.69 ms`
- Memory size: `512 MB`
- Max memory used: `93 MB`

## Java Spring Boot

- Strategy: `aws-serverless-java-container-springboot3`
- Lambda function name: `dmk-java-spring-boot`
- Stack name: `dmk-java-spring-boot`
- HTTP access method: API Gateway HTTP API via SAM `HttpApi`
- Smoke test endpoint: `https://rcd3dhjx0a.execute-api.us-west-2.amazonaws.com/`
- Smoke test result: passed for `/`, `/api/hello/Lan`, and `/api/echo`
- Observed cold start range: `2962.16 ms` to `3558.83 ms`
- Representative cold start: `3558.83 ms`
- Memory size: `1024 MB`
- Max memory used: `183 MB`

## Go Gin

- Strategy: `aws-lambda-go-api-proxy/gin`
- Lambda function name: `dmk-go-gin`
- Stack name: `dmk-go-gin`
- HTTP access method: API Gateway HTTP API via SAM `HttpApi`
- Build result: `sam build --use-container` passed and produced `.aws-sam/build/GinFunction/bootstrap`
- Smoke test endpoint: `https://exl1dzb081.execute-api.us-west-2.amazonaws.com/`
- Smoke test result: passed for `/`, `/api/hello/Lan`, and `/api/echo`
- Observed cold start range: `92.15 ms` to `119.77 ms`
- Representative cold start: `119.77 ms`
- Memory size: `256 MB`
- Max memory used: `30 MB`

## Notes

- AWS credentials are not stored in the repository.
- Terraform was not used.
