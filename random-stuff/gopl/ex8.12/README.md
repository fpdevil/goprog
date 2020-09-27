# Exercise 8.12, Chapter 8

Make the `broadcaster` announce the current set of clients to each new arrival.

This requires that the clients set and the entering and leaving channels record the client name too.

## Build and run the chat server

```shell
⇒  ./chat &
[1] 62844

# later for killing
⇒  pkill chat
[1]  + 62844 terminated  ./chat
```

## Build and run multiple netcat instances

```shell
⇒  go build -o netcat
```

- On Terminal 1
```shell
⇒  ./netcat
You are 127.0.0.1:60740
1 clients online as here: 127.0.0.1:60740

127.0.0.1:60745 has arrived
127.0.0.1:60745: Testing from Terminal 2
Testing from Terminal 1
127.0.0.1:60740: Testing from Terminal 1
^C
```

- On Terminal 2
```shell
⇒  ./netcat
You are 127.0.0.1:60745
2 clients online as here: 127.0.0.1:60740, 127.0.0.1:60745

Testing from Terminal 2
127.0.0.1:60745: Testing from Terminal 2
127.0.0.1:60740: Testing from Terminal 1
127.0.0.1:60740 has left
^C
```