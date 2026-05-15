# AWS Assignment W5

BYOL Lambda migration assignment for Node.js/Express, Java/Spring Boot, and Go/Gin starters.

## Người thực hiện

- `XB-DN26-125`

## Projects

- `xbrain-w5-byol-node-express`
- `xbrain-w5-byol-java-spring-boot`
- `xbrain-w5-byol-go-gin`

## Deployed/Implemented

- Node.js Express was adapted to Lambda using `serverless-http` and SAM HTTP API.
- Java Spring Boot was adapted to Lambda using `aws-serverless-java-container-springboot3` and SAM HTTP API.
- Go Gin was adapted to Lambda using `aws-lambda-go-api-proxy/gin` and SAM HTTP API.

## Deployed API Gateway URLs

| Project | Stack | Lambda Function | API Gateway URL |
|---|---|---|---|
| Node.js Express | `dmk-node-express` | `dmk-node-express` | `https://wffmj4daqd.execute-api.us-west-2.amazonaws.com/` |
| Java Spring Boot | `dmk-java-spring-boot` | `dmk-java-spring-boot` | `https://rcd3dhjx0a.execute-api.us-west-2.amazonaws.com/` |
| Go Gin | `dmk-go-gin` | `dmk-go-gin` | `https://exl1dzb081.execute-api.us-west-2.amazonaws.com/` |

See each project's `NOTES.md` and `template.yaml` for strategy and handler details.
