## Go graceful exit
This is just an example of a golang HTTP application that exits gracefully on ```SIGTERM``` and ```SIGINT``` signals.
This is essential for a true zero downtime especially when your applications run in a docker container.


#### To build
This requires go 1.8+ version and [mux](https://github.com/gorilla/mux).

```bash
$ go get go get github.com/gorilla/mux
```

```bash
$ cd go-graceful-exit
$ go build .
```

#### Run
```bash
$ ./go-graceful-exit
2019/07/15 13:14:29 HTTP server started
```

You can send ```SIGINT``` by just hitting ctrl+c. While ```kill``` command for ```SIGTERM``` signal.
