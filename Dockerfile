# stage -1 : Build
FROM golang:1.23.4-alpine3.21 AS build
WORKDIR /app
COPY . .
RUN go build -o main main.go
#Add curl 
RUN apk add curl
#Download migration using curl
RUN  curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz | tar xvz


# stage-2 : Run stage
FROM alpine:3.21
WORKDIR /app
COPY --from=build /app/main . 
COPY --from=build /app/app.env . 
COPY --from=build /app/migrate ./migrate
COPY db/migration ./migration
COPY start.sh .
COPY wait-for.sh .
EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ] # it will same as running ENTRYPOINT [ "/app/start.sh" , "/app/main"]