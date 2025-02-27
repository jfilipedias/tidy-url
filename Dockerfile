FROM golang:1.24-alpine

ENV API_PORT=
ENV API_ENVIRONMENT=
ENV MONGODB_URI=

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .

EXPOSE ${API_PORT}

CMD [ "go", "run", "./cmd/app", "-port=${API_PORT}}", "-env=${API_ENVIRONMENT}", "-mongo-uri=${MONGODB_URI}" ]