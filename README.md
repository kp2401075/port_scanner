# port_scanner
This is a simple port scanner I wrote for understanding how Go routines work

it is fully functioning so feel free to use it

it takes either 2 or 3 arginemts


## 2 Arguments
1: Host
2: port
``` 
./portscanner localhost 80
``` 
will check wether port 80 is open on the local machine and will show following if open
```
Opened localhost:22
```

## 3 Arguments
1: Host
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