package api

import (
	"encoding/json"
	"fmt"
	"github.com/Anrop/Anrop-Squad/db"
	"net/http"
	"os"
)

// Arma1JSONHandler handles request for Arma1 as JSON
func Arma1JSONHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetArma1Users()
	internalJSONHandler(w, r, users, err)
}

// Arma2JSONHandler handles request for Arma2 as JSON
func Arma2JSONHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetArma2Users()
	internalJSONHandler(w, r, users, err)
}

// Arma3JSONHandler handles request for Arma3 as JSON
func Arma3JSONHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetArma3Users()
	internalJSONHandler(w, r, users, err)
}

// OfpJSONHandler handles request for OFP as JSON
func OfpJSONHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetOfpUsers()
	internalJSONHandler(w, r, users, err)
}

func internalJSONHandler(w http.ResponseWriter, r *http.Request, users *[]db.User, err error) {
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(users)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling json: %q\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
