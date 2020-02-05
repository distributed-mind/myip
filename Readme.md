# My IP

[This](https://github.com/distributed-mind/myip) is a simple golang module which queries DNS for your requesting IP address.
It is much like the web-based services you find when searching for "what is my ip address?" in google. It is based on the information found [here](https://unix.stackexchange.com/a/495954/48033).

## Example

The returned error will report a "network is unreachable" error if either ipv4 or ipv6 is unavailable.
<br />
It will return a "fatal" if both ipv4 and ipv6 are unavailable.

```golang
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
```

### output from ^example
```bash
$ go run example/main.go 
IP: pub.lic.ip.address
Error: ipv6 error: lookup o-o.myaddr.l.google.com on 1.1.1.1:53: dial udp6 [2001:4860:4802:32::a]:53: connect: network is unreachable
$
```

## TODO
- LAN/internal IP discovery
- LAN default gateway discovery

## Contributing
Please submit [an issue](https://github.com/distributed-mind/myip/issues/new) or [pull request](https://github.com/distributed-mind/myip/pull/new/master) on github.
