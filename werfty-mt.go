package main

import "github.com/sparrc/go-ping"
import "fmt"
import "os"
import "net"

func main() {
connectivy_check()
}

func connectivy_check() {
pinger, err := ping.NewPinger("www.google.com")
if err != nil {
        fmt.Println("icmp ping failed");
	os.Exit(1)
}
pinger.Count = 3
pinger.Run() // blocks until finished
//stats := pinger.Statistics() // get send/receive/rtt stats
//fmt.Println(stats)

conn, err := net.Dial("tcp", "google.com:80")
if err != nil {
	fmt.Println("error opening google.com tcp port 80")
	os.Exit(1)
} else {
fmt.Fprintf(conn, "HEAD / HTTP/1.0\r\n\r\n")
conn.Close()
}
}


