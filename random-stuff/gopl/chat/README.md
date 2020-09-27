# Chat server

## Build and run the client and server

Build and start the char server

```shell
⇒  go build -o chat
⇒  ./server/chat &
[1] 46728
```

Build the client `netcat` and run as many instances

```shell
⇒  go build -o netcat
```

- Terminal 1
```shell
⇒  ./client/netcat
You are 127.0.0.1:60550
127.0.0.1:60556 has arrived
127.0.0.1:60556: Hello
This is terminal 1
127.0.0.1:60550: This is terminal 1
127.0.0.1:60556: This is Terminal 2
127.0.0.1:60556: I am leaving...
127.0.0.1:60556 has left
^C
```

- Terminal 2
```shell
⇒  ./client/netcat
You are 127.0.0.1:60556
Hello
127.0.0.1:60556: Hello
127.0.0.1:60550: This is terminal 1
This is Terminal 2
127.0.0.1:60556: This is Terminal 2
I am leaving...
127.0.0.1:60556: I am leaving...
^C
```

Finally kill the server `chat` with `pkill chat`