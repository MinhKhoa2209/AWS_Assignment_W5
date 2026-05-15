# Deployment Notes

## Strategy

I chose `aws-serverless-java-container-springboot3`, the README's canonical Java strategy.

- `Application.java` remains Lambda-unaware.
- `HelloController.java` remains Lambda-unaware.
- `StreamLambdaHandler` is the only Lambda-specific Java entrypoint.
- `maven-shade-plugin` creates a flat Lambda-loadable JAR instead of Spring Boot's nested JAR layout.
- The SAM template keeps the required HTTP API access method and points the Lambda handler at the stream handler class.
