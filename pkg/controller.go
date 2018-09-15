package cloudtables

import (
	"fmt"
	"net/http"
)

type staticController struct{}

func (sc staticController) registerRoutes() {
	http.HandleFunc("/", sc.HandleIndex)
}

var vueController staticController

// BEGIN ROUTES
// Configure route handlers in this section
func (sc staticController) HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// END ROUTES

// Run bootstraps the http server.
func Run(addr string) {
	vueController.registerRoutes()

	css := http.Dir("../../ui/css")
	js := http.Dir("../../ui/js")
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(css)))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(js)))

	http.ListenAndServe(addr, nil)
	fmt.Printf("CloudTables is listening on port %v", addr)
}
