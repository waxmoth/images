package main

import (
	"image-functions/src/routers"
)

func main() {
	r := routers.Routers()
	r.Run(":8080")
}
