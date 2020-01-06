package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	argsWithoutProg := os.Args[1:]
	// host being assigned from command line agument
	host := argsWithoutProg[0]
	if len(argsWithoutProg) == 3 {
		// maximum port number being passed from the command line arguments
		max, _ := strconv.Atoi(argsWithoutProg[2])
		// minimum port number being passed from the command line arguments
		min, _ := strconv.Atoi(argsWithoutProg[1])
		for i := min - 1; i <= max; i++ {
			go rawConnectSingle(host + ":" + strconv.Itoa(i))
		}
	} else if len(argsWithoutProg) == 2 {
		rawConnectSingle(host + ":" + argsWithoutProg[1])
	} else {
		fmt.Print("Incorrect number of arguments \n",
			"This Program requiers atleast 2 aguments 1: hostname & 2: port number \n",
			"3 Aguments will make it check for the range of ports \n",
			"1: host, 2: starting port & 3: ending port\n")
	}
	// waiting for go routine to finish
	defer time.Sleep(time.Second * 5)
}

// rawConnectSignal take host name and port combined as a string and see if it it able to conenct to it
func rawConnectSingle(host string) {
	//fmt.Println("Function not implemeted")
	//	fmt.Println("tcp", host, time.Second)
	conn, err := net.DialTimeout("tcp", host, time.Second)
	if err != nil {
		//	fmt.Println("Connecting error:", err)
	}
	if conn != nil {
		defer conn.Close()
		fmt.Println("Opened", host)
	}
}
