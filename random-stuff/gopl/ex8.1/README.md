# Exercise 8.1 Chapter 8

Modify `clock2` to accept a `port number`, abd write a program, `clockwall`, that acts as a client of several clock servers at once, reading the times from each one and displaying the results in a table, akin to the wall of clocks seen in some business offices. If you have access to geographically distributed computers, run instances remotely; otherwise run local instances on different ports with fake time zones.

```shell
$ TZ=US/Eastern ./clock2 -port 8010 &
$ TZ=Asia/Tokyo ./clock2 -port 8020 &
$ TZ=Europe/London ./clock2 -port 8030 &
$ clockwall NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
```

## Running stats from the console

- *clock server*

```shell
# build the server
⇒  go build -o clock

# running the server (start the clocks)
⇒  TZ=Asia/Calcutta ./clock -port 8010 &
⇒  TZ=Europe/Zurich ./clock -port 8020 &
⇒  TZ=America/Chicago ./clock -port 8030 &

# kill the server (clocks)
⇒  pkill clock
[3]  + 48947 terminated  TZ=America/Chicago ./clock -port 8030
[2]  + 48660 terminated  TZ=Europe/Zurich ./clock -port 8020
[1]  + 47569 terminated  TZ=Asia/Calcutta ./clock -port 8010
```

- *clock client*

```shell
# build the client
⇒  go build -o clockwal

# running the client
⇒  ./clockwall Hyderabad,India=localhost:8010 Zurich,Switzerland=localhost:8020 Chicago,US=localhost:8030
Hyderabad,India      :: 03:08:57 Mon Aug 17 2020
Zurich,Switzerland   :: 23:38:57 Sun Aug 16 2020
Chicago,US           :: 16:38:57 Sun Aug 16 2020
Hyderabad,India      :: 03:08:58 Mon Aug 17 2020
Zurich,Switzerland   :: 23:38:58 Sun Aug 16 2020
Chicago,US           :: 16:38:58 Sun Aug 16 2020
Zurich,Switzerland   :: 23:38:59 Sun Aug 16 2020
Chicago,US           :: 16:38:59 Sun Aug 16 2020
Hyderabad,India      :: 03:08:59 Mon Aug 17 2020
Chicago,US           :: 16:39:00 Sun Aug 16 2020
```