# FTP Client and Server in GO

## Server

Start the `FTP` server ##as below

```shell
$ go run server/FTPServer.go

* GO Simple FTP Server *

# To exit, press Ctrl+C
^Csignal: interrupt
```

## Client

Start the client by providing the target host to connect to

```shell
⇒  go run client/FTPClient.go localhost

* Go FTP Client *

# do a directory listing
dir
server
client

# get present working directory
pwd
Current dir"~/sw/programming/go/src/github.com/fpdevil/goprog/random-stuff/general-stuff/ftp"

# change directory
cd
cd <dir>
cd ../
CD " ../ "

# present working directory
pwd
Current dir"~/sw/programming/go/src/github.com/fpdevil/goprog/random-stuff/general-stuff"

# get directory contents
dir
ftp
binary-tree
html-parsing
sorting
word-frequencies
clock-srv

# exit from the ftp client
quit
```
