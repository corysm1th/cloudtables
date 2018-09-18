package cloudtables

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

// Config holds information about the environment
type Config struct {
	Addr       string `env:"ADDRESS"`
	CertFile   string `env:"CERT_FILE"`
	KeyFile    string `env:"KEY_FILE"`
	CAFile     string `env:"CA_FILE"`
	MutualAuth bool   `env:"MUTUAL_AUTH"`
}

type staticController struct{}

func (sc staticController) registerRoutes() {
	http.HandleFunc("/", sc.HandleIndex)
	http.HandleFunc("api/v1/objects", sc.HandleObjects)
}

var vueController staticController

// BEGIN ROUTES
// Configure route handlers in this section
func (sc staticController) HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (sc staticController) HandleObjects(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// END ROUTES

// Run bootstraps the http server.
func Run(config *Config) {
	vueController.registerRoutes()

	css := http.Dir("../../ui/css")
	js := http.Dir("../../ui/js")
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(css)))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(js)))

	if config.MutualAuth {
		ServeMutualAuth(config)
	} else {
		ServeTLS(config)
	}
}

// ServeTLS listens for TLS connections without client authentication.
func ServeTLS(config *Config) {
	err := http.ListenAndServeTLS(config.Addr, config.CertFile, config.KeyFile, nil)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Error opening TLS listener."))
	}
	fmt.Printf("CloudTables is listening on port %v", config.Addr)
}

// ServeMutualAuth authenticates clients using TLS
func ServeMutualAuth(config *Config) {
	err := http.ListenAndServeTLS(config.Addr, config.CertFile, config.KeyFile, nil)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Error opening TLS listener."))
	}
	fmt.Printf("CloudTables is listening on port %v", config.Addr)
	fmt.Printf("Mutual authentication enabled.")
}
