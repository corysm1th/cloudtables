package cloudtables

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"goji.io"

	assetfs "github.com/elazarl/go-bindata-assetfs"
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
	Storage    string `env:"STORAGE"`
}

// Service holds the cloudtables service resources, like http mux and storage.
type Service struct {
	Router *goji.Mux
	Store  Storage
}

// RegisterRoutes maps API endpoints to Handler functions.
func (svc Service) RegisterRoutes(mux *goji.Mux) {
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
	mux.HandleFunc(pat.Get("/api/v1/objects"), svc.HandleGetObjects)
	mux.HandleFunc(pat.Get("/api/v1/sync"), HandleGetSync)
	mux.HandleFunc(pat.Get("/api/v1/metrics"), HandleGetMetrics)
	mux.HandleFunc(pat.Get("/api/v1/ping"), handlePing) // Healthcheck endpoint
}

// BEGIN ROUTES
// Configure route handlers in this section

// HandleRoot redirects to the UI
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/ui/", http.StatusPermanentRedirect)
}

// HandleGetObjects returns all cloud objects in the database.
func (svc Service) HandleGetObjects(w http.ResponseWriter, r *http.Request) {
	ddb, err := svc.Store.SelectDynamoDBObj()
	if err != nil || ddb == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	eip, err := svc.Store.SelectEC2EIPObj()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	ec2, err := svc.Store.SelectEC2InstObj()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	s3, err := svc.Store.SelectS3BucketObj()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	data := struct {
		DynamoDBs    []*DynamoDBObj `json:"dynamoDBInstances"`
		Ec2IPs       []*EC2EIPObj   `json:"elasticIPs"`
		Ec2Instances []*EC2InstObj  `json:"ec2Instances"`
		S3Buckets    []*S3BucketObj `json:"s3Buckets"`
	}{
		DynamoDBs:    ddb,
		Ec2IPs:       eip,
		Ec2Instances: ec2,
		S3Buckets:    s3,
	}
	response, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Fprint(w, response)
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

// Healthcheck endpoint
func handlePing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// END ROUTES

// Run accepts a config, and bootstraps the http server.
func Run(config *Config, s Storage, listener net.Listener) {
	// Enable debug logging if DEBUG env variable set to true.
	debug := log.Logger{}
	debug.SetOutput(ioutil.Discard)
	if config.Debug {
		debug.SetOutput(os.Stdout)
	}

	mux := goji.NewMux()
	svc := Service{
		Router: mux,
		Store:  s,
	}
	svc.RegisterRoutes(mux)

	server := http.Server{
		Handler: svc.Router,
	}

	go server.Serve(listener)
}
