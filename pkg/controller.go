package cloudtables

import (
	"fmt"
	"log"
	"net/http"

	"goji.io"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/pkg/errors"
	"goji.io/pat"
)

// Config holds information about the environment
type Config struct {
	Addr       string `env:"ADDRESS"`
	CertFile   string `env:"CERT_FILE"`
	KeyFile    string `env:"KEY_FILE"`
	CAFile     string `env:"CA_FILE"`
	MutualAuth bool   `env:"MUTUAL_AUTH"`
}

// RegisterRoutes maps API endpoints to Handler functions.
func RegisterRoutes(mux *goji.Mux) {
	// Load assets from go-bindata
	css := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "ui/css"}
	js := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "ui/js"}

	mux.Handle(pat.Get("/css/"), http.StripPrefix("/css/", http.FileServer(css)))
	mux.Handle(pat.Get("/js/"), http.StripPrefix("/js/", http.FileServer(js)))

	mux.HandleFunc(pat.Get("/"), HandleIndex)

	// API Routes
	mux.HandleFunc(pat.Get("api/v1/objects"), HandleObjects)
}

// BEGIN ROUTES
// Configure route handlers in this section

// HandleIndex serves the main UI via an index page.
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request from: ", r.RemoteAddr, " at ", r.RequestURI)
	buf, err := uiIndexHtml()
	if err != nil {
		http.NotFound(w, r)
	}
	w.Write(buf.bytes)
}

// HandleObjects returns all cloud objects in the database.
func HandleObjects(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// END ROUTES

// Run bootstraps the http server.
func Run(config *Config) {
	mux := goji.NewMux()
	RegisterRoutes(mux)

	if config.MutualAuth {
		err := http.ListenAndServeTLS(config.Addr, config.CertFile, config.KeyFile, mux)
		if err != nil {
			log.Fatalln(errors.Wrap(err, "Error opening TLS listener."))
		}
		fmt.Printf("CloudTables is listening on port %v", config.Addr)
		fmt.Printf("Mutual authentication enabled.")
	} else {
		err := http.ListenAndServeTLS(config.Addr, config.CertFile, config.KeyFile, mux)
		if err != nil {
			log.Fatalln(errors.Wrap(err, "Error opening TLS listener."))
		}
		fmt.Printf("CloudTables is listening on port %v", config.Addr)
	}
}
