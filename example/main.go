// SPDX-License-Identifier: MIT-0
// LICENSE: https://spdx.org/licenses/MIT-0.html

package main

import (
	"fmt"
	"github.com/distributed-mind/myip"
)

func main() {
	ips, err := myip.PublicIPs()
	for _, ip := range ips {
		if len(ip) != 0 {
			fmt.Printf("IP: %v\n", ip.String())
		}
	}
	fmt.Printf("Error: %v\n", err)
}
