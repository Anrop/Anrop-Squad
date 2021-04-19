package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Anrop/Anrop-Squad/api"
	"github.com/Anrop/Anrop-Squad/db"
	"github.com/Anrop/Anrop-Squad/newrelic"
	"github.com/Anrop/Anrop-Squad/static"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")
	databaseURL := os.Getenv("DATABASE_URL")
	newRelicLicenseKey := os.Getenv("NEW_RELIC_LICENSE_KEY")

	if port == "" {
		port = "8080"
	}

	var err error

	statics, err := static.LoadStatics()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error with static files: %q\n", err)
		os.Exit(1)
	}

	err = db.Connect(databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to database: %q\n", err)
		os.Exit(1)
	}

	r := mux.NewRouter()

	if newRelicLicenseKey != "" {
		newrelic.SetupNewRelic(newRelicLicenseKey, r)
	}

	r.HandleFunc("/arma1.json", api.Arma1JSONHandler)
	r.HandleFunc("/arma1.xml", api.Arma1XMLHandler(statics))
	r.HandleFunc("/arma2.json", api.Arma2JSONHandler)
	r.HandleFunc("/arma2.xml", api.Arma2XMLHandler(statics))
	r.HandleFunc("/arma3.json", api.Arma3JSONHandler)
	r.HandleFunc("/arma3.xml", api.Arma3XMLHandler(statics))
	r.HandleFunc("/ofp.json", api.OfpJSONHandler)
	r.HandleFunc("/ofp.xml", api.OfpXMLHandler(statics))
	r.PathPrefix("/").Handler(http.FileServer(http.FS(statics.FS)))

	var handler http.Handler
	handler = handlers.CORS()(r)
	handler = handlers.CompressHandler(handler)

	fmt.Fprintf(os.Stdout, "Server launching on port %s\n", port)

	http.ListenAndServe(":"+port, handler)
}
