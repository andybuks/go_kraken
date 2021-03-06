package main

import (
	"log"

	ws "github.com/aopoltorzhicky/go_kraken/websocket"
)

func main() {
	c := ws.New()
	err := c.Connect()
	if err != nil {
		log.Fatal("Error connecting to web socket : ", err)
	}

	err = c.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// subscribe to BTCUSD, XLMUSD, ADACAD spread
	err = c.SubscribeSpread([]string{ws.ADABTC, ws.XTZBTC, ws.XLMBTC})
	if err != nil {
		log.Fatal(err)
	}

	for obj := range c.Listen() {
		switch obj.(type) {
		case error:
			log.Printf("channel closed: %s", obj)
		case ws.DataUpdate:
			data := obj.(ws.DataUpdate)
			log.Printf("MSG RECV: %#v", data)
		default:
		}
	}
}
