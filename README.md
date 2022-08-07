# blog-go

## How to run?
```
$ git clone <this-repository>
$ cd blog-go
$ make build  -> build all required images
$ kubectl apply -f infra/deploy/on-premise  -> deploy to k8s
```

## How to check deployed pods?
```
$ kubectl -n cshimegi-blog-local get po
```


## APIs
### User API
```
$ curl -v http://localhost/api/users
```

### Post API
```
$ curl -v http://localhost/api/posts
```
