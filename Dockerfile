FROM golang:1.24-alpine

ENV API_PORT=
ENV API_ENVIRONMENT=
ENV MONGODB_URI=

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

EXPOSE $API_PORT

CMD ["sh", "-c", "go tool air -- -port=$API_PORT -env=$API_ENVIRONMENT -db-uri=$MONGODB_URI" ]