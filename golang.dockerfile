FROM golang:1.18-alpine AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /app/start

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/start .

CMD ["./start"]
