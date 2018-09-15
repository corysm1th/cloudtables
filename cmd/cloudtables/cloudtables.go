package main

import (
	"github.com/corysm1th/cloudtables/pkg"
)

const (
	port = ":8443"
)

func main() {
	cloudtables.Run(port)
}
