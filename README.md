## Hello gin and swagger

Try some [gin api examples](https://github.com/gin-gonic/gin#api-examples) and [swagger](https://github.com/swaggo/gin-swagger).

### Download Swag for Go by using:

```
$ go get -u github.com/swaggo/swag/cmd/swag
```
### Add go's bin path to enviroment

```
$ export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
```

### Generate document

```
$ swag init
```

### Run server

```
$ go run .
```

run server again, and then visit http://localhost:8080/swagger/index.html .

## Reference

- [gin](https://github.com/gin-gonic/gin)
- [gin-swagger](https://github.com/swaggo/gin-swagger)
- [swaggo example] (https://github.com/swaggo/swag/tree/master/example)
- [Use Custom Package](https://blog.francium.tech/go-modules-go-project-set-up-without-gopath-1ae601a4e868)
