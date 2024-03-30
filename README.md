# blog-go
A backend API service built with Golang GIN framework for simple blog system.
It's a practice project for learning how to utilize the framework and other libraries to build
a backend system.

## How to run?
```shell
$ git clone <this-repository>
$ cd /project/top

# Build API service image with monolithic structure
$ make build-mono GO_VERSION=<ver>

# Run containers (API/MySQL) locally
$ cd intra/deploy
$ docker-compose up -d
```


## How to migrate database
```shell
# DSN => user:pass@tcp(host:port)/dbname

./migrate -source file://./migrations -database "mysql://$DB_USER:$DB_PASS@tcp($DB_HOST:$DB_PORT)/$DB_NAME" up
```

## APIs
```shell
# User API Endpoint
$ curl -v http://127.0.0.1:<port>/api/users

# Post API Endpoint
$ curl -v http://127.0.0.1:<port>/api/posts
```

## Reference
- [golang-migrate usage](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#usage)
