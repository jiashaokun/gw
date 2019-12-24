package main

import (
	"gw/route"
)

func main() {
	r := route.Route()
	r.Run(":1323")
}
