package cloudtables

import (
	"fmt"
	"net/http"
)

// TLSConfig stores the file paths for TLS certificates
type TLSConfig struct {
	APICert         string
	APIKey          string
	APIAuthority    string
	ClientCert      string
	ClientKey       string
	ClientAuthority string
	UICert          string
	UIKey           string
	UIAuthority     string
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
func Run(addr string, tlsConfig TLSConfig) {
	vueController.registerRoutes()

	css := http.Dir("../../ui/css")
	js := http.Dir("../../ui/js")
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(css)))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(js)))

	// TODO: Read TLSConfig from toml

	// TODO: Check filesystem for UI Certificate

	// TODO: Call Serve or ServeTLS

}

// ServeTLS listens for TLS connections without client authentication.
func ServeTLS(addr string) {
	http.ListenAndServe(addr, nil)
	fmt.Printf("CloudTables is listening on port %v", addr)
}

// ServeMutualAuth authenticates clients using TLS
func ServeMutualAuth(addr string, tlsConfig TLSConfig) {

}
