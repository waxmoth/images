package main

import (
	"image-functions/src/routers"
)

func main() {
	r := routers.Routers()
	r.ForwardedByClientIP = true
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
