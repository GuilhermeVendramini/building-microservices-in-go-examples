# Go kit custom example

http://gokit.io/examples/stringsvc.html

# How to use

```sh
$ go run *.go
```

The following example always will return "Guilherme"
```sh
$ curl -XPOST -d'{"s":"test"}' localhost:8080/get
```

The following example always will return "true"
```sh
$ curl -XPOST -d'{"s":"1"}' localhost:8080/delete
```