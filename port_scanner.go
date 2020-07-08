package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) == 3 {
		// Grabbing hostname/IP from arguments
		host := args[0]
		// Grab min and max from the arguments
		max, _ := strconv.Atoi(args[2])
		min, _ := strconv.Atoi(args[1])
		// Wait group for concurrency
		var wg sync.WaitGroup
		wg.Add(1)

		rawConnectMulti(&wg, host, min, max)
		wg.Wait()

	} else if len(args) == 2 {
		// Grabbing hostname/IP from arguments
		host := args[0]
		// ch := make(chan string)

		rawConnectSingle(host + ":" + args[1])
	} else {
		// Output a bit of help on user error
		fmt.Print("Incorrect number of arguments \n",
			"This Program works with wither 2 or 3 aguments \n",
			"2 Arguments mode :\n",
			"	1: hostname \n",
			" 	2: port number \n",
			"3 Aguments mode: \n",
			"	1: host \n",
			"	2: starting port\n ",
			"	3: ending port\n")
	}

}

// rawConnectMulti will loop over port range, It also take waitgroup pointer that will wait for it to finish
func rawConnectMulti(wg *sync.WaitGroup, host string, start int, stop int) {

	// looping from min to max of the range
	for i := start - 1; i <= stop; i++ {
		go rawConnectSingle(host + ":" + strconv.Itoa(i))

	}
	fmt.Println("Done")
	wg.Done()

}

//rawConnectSingle will try to connect a host on given port
func rawConnectSingle(host string) {
	// Dialing connection to host:port
	conn, _ := net.DialTimeout("tcp", host, time.Second)
	// Print if connection can be opened
	if conn != nil {
		defer conn.Close()
		fmt.Println("Open", host)
	}
}
