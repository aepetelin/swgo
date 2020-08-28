# Go Swagger. API documetation example

* based on [swaggo](https://github.com/swaggo/swag) and [swaggo/http-swagger](https://github.com/swaggo/http-swagger)
* public and private endpoints 
* authorization by token in header

Download Swag for Go by using:

```
go get github.com/swaggo/swag/cmd/swag
```

Init in the project root:

```
swag init --parseDependency
go build
./swgo 
```

Open http://localhost:1323/swagger/index.html

(c) MIT Licence