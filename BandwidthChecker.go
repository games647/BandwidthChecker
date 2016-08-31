package main

import (
	"log"
	"net/url"
	"github.com/huin/goupnp/dcps/internetgateway1"
)

const ROUTER_IP = "http://fritz.box:49000/igddesc.xml"

func main() {
	gateway2()
}

func gateway2() {
	routerPath, err := url.Parse(ROUTER_IP)
	if err != nil {
		log.Fatal(err)
	}

	clients, err := internetgateway1.NewWANCommonInterfaceConfig1ClientsByURL(routerPath)
	if err != nil {
		log.Fatal("Error discovering service with UPnP:", err)
	}

	for _, client := range clients {
		if recv, err := client.GetTotalBytesReceived(); err != nil {
			log.Println("Error requesting bytes received:", err)
		} else {
			log.Println("Bytes received:", recv)
		}

		if sent, err := client.GetTotalBytesSent(); err != nil {
			log.Println("Error requesting bytes sent:", err)
		} else {
			log.Println("Bytes sent:", sent)
		}
	}
}