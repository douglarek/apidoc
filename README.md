# ApiDoc

Auto generate api docs from golang packages.

Example:

```
$ go run *.go -p . -t 'http.ResponseWriter,*http.Request' # a sample for http handler func; multiple parameters splited by `,`.
```
