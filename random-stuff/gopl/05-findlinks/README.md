# HTML Parsing

The Chapter 05 discusses about functions and there is a discussion about using `go` html parser for handling html documents. Specifically, `html` from a link is parsed and all the links are printed.

The go package `golang.org/x/net/html` is used for parsing.

## Input

The input to the `main` was originally output from the `fetch` example of chapter `01`. But here a separate function to get the html data from specified input link as argument is within the code under `fetch` function.

So usage becomes:

```bash
⇒  go run main.go https://golang.org
https://support.eji.org/give/153413/#!/donation/checkout
/
/doc/
/pkg/
/project/
/help/
/blog/
https://play.golang.org/
/dl/
https://tour.golang.org/
https://blog.golang.org/
/doc/copyright.html
/doc/tos.html
http://www.google.com/intl/en/policies/privacy/
http://golang.org/issues/new?title=x/website:
https://google.com
```

Library reference avaialble [at]('https://pkg.go.dev/golang.org/x/net/html?tab=doc')

Mozilla html `DOM` [reference](https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model) will be usedulf while parsing