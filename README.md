## Hello gin and swagger

Try some [API examples](https://github.com/gin-gonic/gin#api-examples) and add swagger.

### Run server

```
$ go run .
```
### Download Swag for Go by using:

```
$ go get -u github.com/swaggo/swag/cmd/swag
```
### Generate document

```
$ swag init
```
run server again, and then visit http://localhost:8080/swagger/index.html .

## Reference

- [gin](https://github.com/gin-gonic/gin)
- [Use Custom Package](https://blog.francium.tech/go-modules-go-project-set-up-without-gopath-1ae601a4e868)
- [gin-swagger](https://github.com/swaggo/gin-swagger)
