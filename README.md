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

## API Calls with using docker-compose
- For ARM machine, you can use colima like this to ensure the env setup is workable for other arch machines
  - **colima start -p x86 --arch x86_64 --with-kubernetes**
  - If you use **--network-address** option as well, you will have to use ip address (127.0.0.1) instead
```shell
# User API Endpoint
$ curl -v http://localhost:<port>/api/users

# Post API Endpoint
$ curl -v http://localhost:<port>/api/posts

# Health API Endpoint
$ curl -v http://localhost:<port>/api/health
$ curl -v http://localhost:<port>/api/health/all
```

## API Calls with using k3s
- nginx ingress controller is required
```shell
# User API Endpoint
$ curl -v http://localhost/api/users

# Post API Endpoint
$ curl -v http://localhost/api/posts

# Health API Endpoint
$ curl -v http://localhost/api/health
$ curl -v http://localhost/api/health/all
```

## Reference
- [golang-migrate usage](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#usage)
- [nginx ingress controller on k3s](https://medium.com/@alesson.viana/installing-the-nginx-ingress-controller-on-k3s-df2c68cae3c8)
- [colima localhost/ip address](https://github.com/abiosoft/colima/issues/562#issuecomment-1371331348)
