# Echo Server over TCP

A Simple TCP Server to echo back the request it receives

## Sample run

- Server
```shell
⇒  go run main.go
2020/10/18 14:03:45 Listening on 0.0.0.0:8080
2020/10/18 14:05:15 Received connection
2020/10/18 14:05:30 Received 20 bytes: test of echo server

2020/10/18 14:05:30 Writing data
2020/10/18 14:05:47 Client disconnected
^Csignal: interrupt
```

- Client
```shell
⇒  nc -v localhost 8080
Connection to localhost port 8080 [tcp/http-alt] succeeded!
test of echo server
test of echo server
^C
```
