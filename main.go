package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Anrop/Anrop-Squad/api"
	"github.com/Anrop/Anrop-Squad/db"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")
	databaseURL := os.Getenv("DATABASE_URL")

	if port == "" {
		port = "8080"
	}

	var err error

	err = db.Connect(databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to database: %q\n", err)
		os.Exit(1)
	}

	r := mux.NewRouter()
	r.HandleFunc("/arma1.json", api.Arma1JSONHandler)
	r.HandleFunc("/arma1.xml", api.Arma1XMLHandler)
	r.HandleFunc("/arma2.json", api.Arma2JSONHandler)
	r.HandleFunc("/arma2.xml", api.Arma2XMLHandler)
	r.HandleFunc("/arma3.json", api.Arma3JSONHandler)
	r.HandleFunc("/arma3.xml", api.Arma3XMLHandler)
	r.HandleFunc("/ofp.json", api.OfpJSONHandler)
	r.HandleFunc("/ofp.xml", api.OfpXMLHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	var handler http.Handler
	handler = handlers.CORS()(r)
	handler = handlers.CompressHandler(handler)

	fmt.Fprintf(os.Stdout, "Server launching on port %s\n", port)

	http.ListenAndServe(":"+port, handler)
}
