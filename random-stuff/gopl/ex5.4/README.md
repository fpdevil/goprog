# Exercise 5.4

Extend the `visit` function so that it extracts other kinds of links from the document, such as *images*, *scripts*, and *style sheets*.

## Usage

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

Library reference available [at]('https://pkg.go.dev/golang.org/x/net/html?tab=doc')

Mozilla html `DOM` [reference](https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model) will be usedulf while parsing