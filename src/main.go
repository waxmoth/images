package main

import (
	"image-functions/src/routers"
	"log"
)

func main() {
	r := routers.Routers()
	r.ForwardedByClientIP = true
	if r.SetTrustedProxies([]string{"127.0.0.1"}) != nil {
		log.Fatal("Failed set trusted proxies")
	}
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
