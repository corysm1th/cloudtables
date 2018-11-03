package main

import (
	"crypto/tls"
	"log"
	"net"

	"github.com/caarlos0/env"
	"github.com/corysm1th/cloudtables/pkg"
	"github.com/pkg/errors"
)

// Set some default values for the environment
var (
	config = cloudtables.Config{
		Addr:       ":8443",
		CertFile:   "tls/cert.pem",
		KeyFile:    "tls/cert-key.pem",
		CAFile:     "tls/ca.pem",
		MutualAuth: false,
		Storage:    "memory",
	}
)

func provisionStorage(env string) cloudtables.Storage {
	switch env {
	case "memory":
		return cloudtables.NewStorageMem()
	default:
		log.Printf("No storage type %s", config.Storage)
		log.Println("Defaulting to in-memory storage.")
		return cloudtables.NewStorageMem()
	}
}

func provisionTLS(config *cloudtables.Config) (net.Listener, error) {
	certs, err := tls.LoadX509KeyPair(config.CertFile, config.KeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to load TLS pair")
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{certs},
	}
	listener, err := tls.Listen("tcp", config.Addr, tlsConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to listen on %v", config.Addr)
	}

	return listener, nil
}

func main() {
	// Parse environment variables
	env.Parse(&config)

	// Provision Storage
	storage := provisionStorage(config.Storage)

	// Initialize TLS listener
	listener, err := provisionTLS(&config)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the state machine
	state := cloudtables.NewState()

	cloudtables.Run(&config, storage, listener, state)
}
