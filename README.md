# heroku-test

## Run

Start Redis:

```sh
docker run --name some-redis -p 6379:6379 -d redis
```

Start server:

```sh
PORT=5000 REDIS_URL=redis://localhost:6379/0 go run server/server.go
```

## Build

```sh
go build -o bin/server server/server.go
```
