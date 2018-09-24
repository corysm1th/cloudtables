package cloudtables

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
	Debug      bool   `env:"DEBUG"`
}

// RegisterRoutes maps API endpoints to Handler functions.
func RegisterRoutes(mux *goji.Mux) {
	// Load assets from go-bindata
	css := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "ui/css"}
	js := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "ui/js"}
	index := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "ui"}

	mux.Handle(pat.Get("/css/*"), http.StripPrefix("/css/", http.FileServer(css)))
	mux.Handle(pat.Get("/js/*"), http.StripPrefix("/js/", http.FileServer(js)))
	mux.Handle(pat.Get("/ui/*"), http.StripPrefix("/ui/", http.FileServer(index)))
	mux.HandleFunc(pat.Get("/"), HandleRoot)
	mux.HandleFunc(pat.Get("//"), HandleRoot)

	// API Routes
	mux.HandleFunc(pat.Get("/api/v1/objects"), HandleGetObjects)
	mux.HandleFunc(pat.Get("/api/v1/sync"), HandleGetSync)
	mux.HandleFunc(pat.Get("/api/v1/metrics"), HandleGetMetrics)
}

// BEGIN ROUTES
// Configure route handlers in this section

// HandleRoot redirects to the UI
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/ui/", http.StatusPermanentRedirect)
}

// HandleGetObjects returns all cloud objects in the database.
func HandleGetObjects(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// HandleGetSync triggers a sync action between the cloud providers
// and the local database
func HandleGetSync(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}

// HandleGetMetrics retrieves metrics about the objects stored in cloudtables
func HandleGetMetrics(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// END ROUTES

// Run accepts a config, and bootstraps the http server.
func Run(config *Config) {
	// Enable debug logging if DEBUG env variable set to true.
	debug := log.Logger{}
	debug.SetOutput(ioutil.Discard)
	if config.Debug {
		debug.SetOutput(os.Stdout)
	}

	mux := goji.NewMux()
	RegisterRoutes(mux)

	if config.MutualAuth {
		err := http.ListenAndServeTLS(config.Addr, config.CertFile, config.KeyFile, mux)
		if err != nil {
			log.Fatalln(errors.Wrap(err, "Error opening TLS listener."))
		}
		log.Printf("CloudTables is listening on port %v", config.Addr)
		log.Printf("Mutual authentication enabled.")
	} else {
		err := http.ListenAndServeTLS(config.Addr, config.CertFile, config.KeyFile, mux)
		if err != nil {
			log.Fatalln(errors.Wrap(err, "Error opening TLS listener."))
		}
		log.Printf("CloudTables is listening on port %v", config.Addr)
	}
}
