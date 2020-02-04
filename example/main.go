package main

import (
	"fmt"
	"github.com/distributed-mind/myip"
)

func main() {
	ips, _ := myip.PublicIPs()
	for _, ip := range ips {
		fmt.Printf("%v\n", ip.String())
	}
}
