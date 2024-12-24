# stage -1 : Build
FROM golang:1.23.4-alpine3.21 AS build
WORKDIR /app
COPY . .
RUN go build -o main main.go

# stage-2 : Run stage
FROM alpine:3.21
WORKDIR /app
COPY --from=build /app/main . 
COPY --from=build /app/app.env . 
EXPOSE 8080
CMD [ "/app/main" ]