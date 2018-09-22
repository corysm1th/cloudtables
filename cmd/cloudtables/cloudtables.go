package main

import (
	"github.com/caarlos0/env"
	"github.com/corysm1th/cloudtables/pkg"
)

// Set some default values for the environment
var (
	config = cloudtables.Config{
		Addr:       ":8443",
		CertFile:   "tls/cert.pem",
		KeyFile:    "tls/cert-key.pem",
		CAFile:     "tls/ca.pem",
		MutualAuth: false,
	}
)

func main() {
	env.Parse(&config)
	cloudtables.Run(&config)
}
