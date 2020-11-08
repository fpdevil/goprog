# Exercise 8.13, Chapter 8

Make the *chat server* disconnect idle _clients_, such as those that have sent no messages in the last *five* minutes.

_Hint: calling `conn.Close()` in another `goroutine` unblocks active `Read` calls such as the one done by `input.Scan()`_.

## Build and run the chat server

Testing with a timeout value of `10` seconds rather than `5` minutes

```shell
⇒  ./chat &
[1] 62844

# later for killing
⇒  pkill chat
[1]  + 62844 terminated  ./chat
```

## Run multiple netcat instances

- On Terminal 1
```shell
⇒  nc localhost 8000
You are 127.0.0.1:57562
1 clients online as here: 127.0.0.1:57562

127.0.0.1:57564 has joiined
Test from 1
Message from 127.0.0.1:57562: Test from 1
Message from 127.0.0.1:57564: Test from 2
```

- On Terminal 2
```shell
⇒  nc localhost 8000
You are 127.0.0.1:57564
2 clients online as here: 127.0.0.1:57562, 127.0.0.1:57564

Message from 127.0.0.1:57562: Test from 1
Test from 2
Message from 127.0.0.1:57564: Test from 2
disconnecting 127.0.0.1:57562 as idle for 10s seconds
127.0.0.1:57562 has left
```