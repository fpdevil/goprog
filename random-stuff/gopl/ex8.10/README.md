# Exercise 8.10 Go Programming Language

`HTTP` requests may be canceled by closing the optional `Cancel` channel in the `http.Request` struct.
Modify the **web crawler** of section *8.6* to support cancellation.

*Hint:* the `http.Get` convenience function does not give you an opportunity to customize a Request. Instead, create the request using `http.NewRequest`, set its Cancel field, then perform the request by calling `http.DefaultClient.Do(req)`.
