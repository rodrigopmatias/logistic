FROM golang as build

WORKDIR /app
COPY . .
RUN go build -o app-server cmd/server/main.go && \
    go build -o app-migrate cmd/migrate/main.go

FROM ubuntu

COPY --from=build /app/app-server /bin/app-server
COPY --from=build /app/app-migrate /bin/app-migrate