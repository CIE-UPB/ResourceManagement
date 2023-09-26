# ResourceManagement

RestApi backend code for the administration area, enabling user management, resource visibility.

## Start the necessary services:

```bash
docker-compose up -d
```

## ENV

The init template in Dockerfile:

``````bash
ENV DATABASE_NAME="cie-db" \
 DATABASE_USER="postgres" \
 DATABASE_PASSWORD="postgres" \
 DATABASE_HOST="postgresql" \
 DATABASE_PORT="5432" \
 DATABASE_SSLMODE="disable" \
 DATABASE_LOGMODE="false" \
 GIN_MODE="debug" \
 SERVER_PORT="8080" \
 ENV="development" \
 LOG_OUTPUT="./server.log" \
 LOG_LEVEL="info" \
 CONTEXT_TIMEOUT=2
``````
