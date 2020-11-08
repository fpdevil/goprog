# Exercise 8.14, Chapter 8

Change the `chat` server's network protocol so that each client provides it's name on entering.
Use that name instead of the network address when performing each message with its senders identity.

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
enter your name: MIckey
You are MIckey
1 clients online as here: MIckey

Donald has joiined
Hey Donald How are you?
Message from MIckey: Hey Donald How are you?
Message from Donald: Good Mickey
```

- On Terminal 2
```shell
⇒  nc localhost 8000
enter your name: Donald
You are Donald
2 clients online as here: MIckey, Donald

Message from MIckey: Hey Donald How are you?
Good Mickey
Message from Donald: Good Mickey
disconnecting 127.0.0.1:57843 as idle for 30s seconds
MIckey has left
```