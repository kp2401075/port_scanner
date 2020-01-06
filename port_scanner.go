package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	// host -> the remote host
	// timeoutSecs -> the timeout value
	host := "pucq-admin-01.networks.in.telstra.com.au"
	// maximum port number
	max := 800
	// minimum port number
	min := 22

	for i := min; i <= max; i++ {

		go rawConnectSingle(host + ":" + strconv.Itoa(min+i))

	}

	time.Sleep(time.Second * 2)

}
func rawConnectSingle(host string) {
	//fmt.Println("Function not implemeted")
	//	fmt.Println("tcp", host, time.Second)
	conn, err := net.DialTimeout("tcp", host, time.Second)
	if err != nil {
		fmt.Println("Connecting error:", err)
	}
	if conn != nil {
		defer conn.Close()
		fmt.Println("Opened", host)
	}
}

// func raw_connect(host string, ports []string) {

// 	for _, port := range ports {
// 		timeout := time.Second
// 		//port_string := strconv.Itoa(port)
// 		//	fmt.Println("tcp", net.JoinHostPort(host, port), timeout)
// 		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
// 		if err != nil {
// 			fmt.Println("Connecting error:", err)
// 		}
// 		if conn != nil {
// 			defer conn.Close()
// 			fmt.Println("Opened", net.JoinHostPort(host, port))
// 		}
// 	}
// }
