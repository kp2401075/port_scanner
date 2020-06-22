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

	//Creating and adding waitgroup for go routine
	var wg sync.WaitGroup
	wg.Add(1)

	// Grabbing hostname/IP from arguments
	host := args[0]

	if len(args) == 3 {
		// Grab min and max from the arguments
		max, _ := strconv.Atoi(args[2])
		min, _ := strconv.Atoi(args[1])

		go func() {
			rawConnectMulti(host, min, max)

			wg.Done()
		}()
	} else if len(args) == 2 {
		rawConnectSingle(host + ":" + args[1])
	} else {
		fmt.Print("Incorrect number of arguments \n",
			"This Program requiers atleast 2 aguments 1: hostname & 2: port number \n",
			"3 Aguments will make it check for the range of ports \n",
			"1: host, 2: starting port & 3: ending port\n")
	}
	// wait before completion of scaning
	wg.Wait()
}

// rawConnectMulti will loop over port range
func rawConnectMulti(host string, start int, stop int) {
	// looping from min to max of the range
	for i := start - 1; i <= stop; i++ {
		rawConnectSingle(host + ":" + strconv.Itoa(i))
	}
}

//rawConnectSingle will try to connect a host on given port
func rawConnectSingle(host string) {

	conn, _ := net.DialTimeout("tcp", host, time.Second)
	// Print if connection can be opened
	if conn != nil {
		defer conn.Close()
		fmt.Println("Opened", host)
	}
}
