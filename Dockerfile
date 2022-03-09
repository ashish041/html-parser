# syntax=docker/dockerfile:1

FROM golang:1.17-alpine as build
WORKDIR /app
COPY go.mod ./
COPY cmd ./cmd/
COPY internal ./internal/
ADD nohup.out ./
RUN go mod tidy
RUN CGO_ENABLED=0 go build cmd/httpserver/main.go

FROM alpine:latest
WORKDIR /app
ADD nohup.out ./
COPY --from=build /app/main ./
EXPOSE 8080
CMD [ "nohup","/app/main","&" ]
