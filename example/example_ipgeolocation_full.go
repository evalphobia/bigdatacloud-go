package main

import (
	"flag"
	"fmt"

	"github.com/evalphobia/bigdatacloud-go/bigdatacloud"
	"github.com/evalphobia/bigdatacloud-go/config"
)

// nolint
func main() {
	var ipaddr string
	flag.StringVar(&ipaddr, "ipaddr", "", "set target ip address")
	flag.Parse()

	conf := config.Config{
		APIKey: "",
		Debug:  true,
	}

	svc, err := bigdatacloud.New(conf)
	if err != nil {
		panic(err)
	}

	resp, err := svc.IPGeolocationFull(ipaddr)
	fmt.Printf("[%+v]\n", resp)
	if err != nil {
		panic(err)
	}
}
