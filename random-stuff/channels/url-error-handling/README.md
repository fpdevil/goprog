# Error Handling while making concurrent calls

Here, we consider a simple concurrent solution using `chnnels` and `goroutines` while making multiple concurrent url's to several urls at the same time.

The main take away here is how the `errors` are handled while invoking the `http` calls to the `url's`.

The errors are considered as _first-class citizens_ when constructing values to return from `goroutines`. If the `goroutine` is able to produce errors, those errors should be tightly couples with our result type and passed along through the same lines of communication, just like regular asynchronous functions.

## Result stub

We construct the below `struct` type to hold the response and the error data.

```go
type Result struct {
	Error    error
	URL      string
	Response *http.Response
}
```

## Stopping criteria

We can consider to enforce a stopping criteria to force close further processing and return back the results till that point if we think that the error rate has increased more than the anticipated.

Here, as an example we consider a numeric limit of `n = 3` to enforce such condition. It keeps track of the error count during each call and once it reaches the threshold value `n`, further processing is stooped and the result is returned.

Here is a sample response...

```bash
⇒  go run .

	Error Handling in concurrent programs
	Handling of errors in a sane way while making concurrent
	requests to multiple urls.


Response from {http://httpbin.org}: "200 OK"
Response from {https://www.rust-lang.org}: "200 OK"
error from {http://1.2.3.4}: Get "http://1.2.3.4": dial tcp 1.2.3.4:80: i/o timeout
error from {a}: Get "a": unsupported protocol scheme ""
error from {b}: Get "b": unsupported protocol scheme ""
Too many errors... breaking the call!
```
