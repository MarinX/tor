# Unofficial Tor source with circuit limiter

## About
This is an unofficial Tor version which contains a circuit limiter.
If there is a multiple connections going through a circuit and hits threashold limit, it will close the connection and circuit, thus the attacker will need to create a new circuit if he wants to DDoS the onion service.



## How it works

In `connection_edge_finished_connecting` function of `connection_edge.c` source, once we have a connection, we give the circuit id to the go handler 

```c
GoCircuitHandler((unsigned)edge_conn->on_circuit->n_circ_id)
```
If a threshold is hit, the error will be returend telling to close the connection
```c
edge_conn->end_reason = END_STREAM_REASON_RESOURCELIMIT;
edge_conn->edge_has_sent_end = 1;
connection_close_immediate(conn);
connection_mark_for_close(conn);
```
This layer protects your onion service for making request and Tor for resource handling (no reading or writing is happening)

## Building

To build, you will need a standard build tools for building a Tor from source and a <a href="https://golang.org/">Golang compiler</a>

### Building a shared library

Before building a Tor, navigate to

```sh
cd src/lib/goddos 
```

Execute `make` to build a shared lib (.so)
```sh
make
```

You should now have the `libgoddos.h` header file and `libgoddos.so` shared object.

### Install a shared library

Depending on your system, copy the shared library.

For example, in MacOS you will copy to
```sh
cp libgoddos.so /usr/local/lib
```

### Build Tor

Follow the official Tor guide on how to build the source and install any missing build tools
```sh
./autogen.sh
./configure
make
make install
```

## Making changes
If you want to make changes about threashold and ban time, open `main.go` and modify this 2 constants
```go
const maxRequestsPerSecond = 3
const banTime = 30 * time.Minute
```
Once you are satified with configuration, execute `make` and replace already installed `libgoddos.so` with new one.

Probably, the best thing is to move configuration for threshold and ban time to `torrc` config file and pass it in `Go` lib. (PR welcome)

**NOTE**

You dont need to install the Tor source again,
since the Tor will load the shared library at runtime.


## Contributing
Feel free to contribute, please check open issues. 