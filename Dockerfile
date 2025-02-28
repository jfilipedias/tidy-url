FROM golang:1.24-alpine

ENV API_PORT=
ENV API_ENVIRONMENT=
ENV MONGODB_URI=

WORKDIR /app
COPY . .
RUN go mod download

EXPOSE $API_PORT

CMD ["sh", "-c", "go tool air" ]