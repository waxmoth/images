package main

import (
	"image-functions/src/routers"
)

func main() {
	r := routers.Routers()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
