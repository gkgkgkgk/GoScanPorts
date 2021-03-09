package main

import (
	"log"
	"net"
	"time"
	"strings"
	"flag"
	"strconv"
)

func main() {
	log.SetFlags(2)

	minPort, maxPort, host, timeout := getInfo()
	log.Printf("%s%d%s%d%s%d", "Scanning ports ", minPort, " through ", maxPort, " on host " + host + " with timeout ", timeout)

	closedPorts := 0
	filteredPorts := 0

	for i := minPort; i <= maxPort; i++ {
		conn, err := net.DialTimeout("tcp", host + ":" + strconv.Itoa(i), time.Duration(timeout) * time.Millisecond)

		if err != nil {
				if strings.HasSuffix(err.Error(), "connection refused"){
					closedPorts++;
				} else if strings.HasSuffix(err.Error(), "i/o timeout"){
					filteredPorts++;
				}
		} else {
			service := getService(i)
			if service != "" {
				log.Printf("%s%d%s","Port ",i," is an open " + service + " port")
			} else {
				log.Printf("%s%d%s","Port ",i," is open.")
			}

			defer conn.Close()
		}
	}

	log.Printf("%d%s", closedPorts, " closed ports detected")
	log.Printf("%d%s", filteredPorts, " filtered ports detected")
}

func getInfo() (int, int, string, int) {
	f := flag.String("p", "0:1024", "Specified Ports")
	t := flag.Int("t", 50, "Specified Timeout")
	h := flag.String("host", "127.0.0.1", "Specified Host")

	flag.Parse()

	ports := strings.Split(*f, ":")

	p1, _ := strconv.Atoi(ports[0])
	p2, _ := strconv.Atoi(ports[0])

	if len(ports) > 1 {
		p2, _ = strconv.Atoi(ports[1])
	}


	return p1, p2, *h, *t
}

func getService(port int) (string){
	switch port {
		case 22:
			return "SSH"
		case 443:
			return "HTTPS"
		case 80:
			return "HTTP"
		case 21:
			return "FTP"
		case 53:
			return "DNS"
		case 23:
			return "TELNET"
		case 25:
			return "SMTP"
		case 143:
			return "IMAP"
		case 110:
			return "POP3"
		case 123:
			return "NTP"
		case 194:
			return "IRC"
		case 3306:
			return "MYSQL"
		case 213:
			return "IPX"
		case 135:
			return "MSRPC"
		case 1723:
			return "PPTP"
		default:
			return ""
	}
}
