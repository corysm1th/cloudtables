package main

import (
	"github.com/corysm1th/cloudtables/pkg"
)

const (
	addr = ":8443"
)

func main() {
	tlsConfig := cloudtables.TLSConfig{}
	cloudtables.Run(addr, tlsConfig)
}
