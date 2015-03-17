// Package main provides ...
package main

import (
	"fmt"
	"github.com/Zuckonit/taobaoip"
	"os"
)

func Help() {
	prog = os.Args[0]
	fmt.Printf(`
%s <IP>\r\n
   `, prog)
}

func main() {
	if len(os.Args) == 1 {
		Help()
	} else {
		ip := os.Args[1]
		req := taobaoip.Req{ip: ip}
		rb, err := req.URLOpen()
		if err != nil || rb == nil {
			return
		}
		rb.Print()
	}
}
