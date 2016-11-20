package main

import "github.com/sparrc/go-ping"
import "flag"
import "fmt"
import "net"
import "os"
import "time"

var verbose *bool

func main() {

	var usage = `
Usage:
    werfty-mt [-v verbose] action [action parameters]

    where action might be
    
    conncheck [target]	check for network connectivity (ICMP Echo ping/TCP Port 80) default target: www.google.com

    Examples: 
    # check if we can reach www.google.com (default target)
    werfty-mt conncheck
    # check if we can reach www.twitter.com
    werfty-mt conncheck www.twitter.com

`
	verbose = flag.Bool("v", false, "")
	flag.Usage = func() {
		fmt.Printf(usage)
	}
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	actions := []string{"conncheck"}
	action := flag.Arg(0)

	if !Include(actions, action) {
		flag.Usage()
		return
	}

	if action == "conncheck" {
		target := ""
		if flag.NArg() == 2 {
			target = flag.Arg(1)
		} else {
			target = "www.google.com"
		}
		connectivy_check(target)
	}
}

func connectivy_check(target string) {
	pinger, err := ping.NewPinger(target)
	if err != nil {
		fmt.Println("ICMP ping to " + target + " failed")
		os.Exit(1)
	}
	pinger.Count = 1
	pinger.Timeout = time.Second * 5

	pinger.OnFinish = func(stats *ping.Statistics) {
		if stats.PacketsRecv == 1 {
			if *verbose {
				fmt.Printf("Echo reply recieved from %s (%s) => ICMP test successful\n", target, stats.Addr)
			}
		} else {
			if *verbose {
				fmt.Println("Did not recieve ICMP reply from " + target + ". ICMP test failed")
			}
		}
	}
	pinger.Run() // blocks until finished

	conn, err := net.Dial("tcp", target+":80")
	if err != nil {
		if *verbose {
			fmt.Println("error opening " + target + " tcp port 80")
		}
		os.Exit(1)
	} else {
		fmt.Fprintf(conn, "HEAD / HTTP/1.0\r\n\r\n")
		conn.Close()
		if *verbose {
			fmt.Println("TCP Port 80 port to " + target + " opened => TCP test successful")
		}
	}

	if *verbose {
		fmt.Println("connection check successful")
	}
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}
