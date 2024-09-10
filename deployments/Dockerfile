FROM golang:1.20.4-alpine3.17
RUN apk add --no-progress --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

ARG SERVICE_NAME
RUN go build -tags musl -ldflags '-extldflags "-static"' -o main ./cmd/${SERVICE_NAME}

ARG SERVICE_PORT
EXPOSE ${SERVICE_PORT}

CMD ["./main"]