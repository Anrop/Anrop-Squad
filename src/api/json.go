package api

import (
	"db"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func Arma1JsonHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetArma1Users()
	internalJsonHandler(w, r, users, err)
}

func Arma2JsonHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetArma2Users()
	internalJsonHandler(w, r, users, err)
}

func Arma3JsonHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetArma3Users()
	internalJsonHandler(w, r, users, err)
}

func OfpJsonHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetOfpUsers()
	internalJsonHandler(w, r, users, err)
}

func internalJsonHandler(w http.ResponseWriter, r *http.Request, users *[]db.User, err error) {
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
