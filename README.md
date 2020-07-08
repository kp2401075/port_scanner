#### Disclaimer:
This program may not completely reliable, This program is not designed or intended for use in the design, construction, operation or maintenance of any nuclear facility.  I disclaims any express or implied  warranty  of  fitness  for such  uses

# What is this?
This is a simple port scanner I wrote for understanding how Go routines work. Feel free to compile and use it. It takes either 2 or 3 arguments


## 2 Arguments

1: Hostname or IP address

2: port number
``` 
./portscanner localhost 80
``` 
will check wether port 80 is open on the local machine and will show following if open
```
Opened localhost:80
```

## 3 Arguments
1: Hostname or IP address

2: Starting port

3: Ending port
``` 
./portscanner localhost 10 10000
``` 
will check wether ports between 10 and 10000 are open on the local machine and will show following for the open ports
```
Opened localhost:22
Opened localhost:631
```

## Compilation
If you have make build system installed just run
```
make clean
make
```
else you can run following to create a binary for your package
```
go build port_scanner.go
```