// Package main provides ...
package main

import (
	"fmt"
	"os"
	"path"
	"taobaoip"
)

func Help() {
	prog := path.Base(os.Args[0])
	fmt.Printf("%s <IP>\r\n", prog)
}

func main() {
	if len(os.Args) == 1 {
		Help()
		os.Exit(1)
	} else {
		ip := os.Args[1]
		req := taobaoip.Req{IP: ip}
		rb, err := req.URLOpen()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}
		rb.Print()
	}
}
