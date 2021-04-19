package api

import (
	"encoding/xml"
	"fmt"
	"github.com/Anrop/Anrop-Squad/db"
	"github.com/Anrop/Anrop-Squad/static"
	"net/http"
	"os"
)

// Arma1XMLHandler handles requests for Arma1 as XML
func Arma1XMLHandler(statics *static.Statics) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := db.GetArma1Users()
		internalXMLHandler(w, r, statics, *users, err)
	}
}

// Arma2XMLHandler handles requests for Arma2 as XML
func Arma2XMLHandler(statics *static.Statics) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := db.GetArma2Users()
		internalXMLHandler(w, r, statics, *users, err)
	}
}

// Arma3XMLHandler handles requests for Arma3 as XML
func Arma3XMLHandler(statics *static.Statics) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := db.GetArma3Users()
		internalXMLHandler(w, r, statics, *users, err)
	}
}

// OfpXMLHandler handles requests for OFP as XML
func OfpXMLHandler(statics *static.Statics) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := db.GetOfpUsers()
		internalXMLHandler(w, r, statics, *users, err)
	}
}

func internalXMLHandler(w http.ResponseWriter, r *http.Request, statics *static.Statics, users []db.User, err error) {
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	doctype := fmt.Sprintf("<!DOCTYPE squad SYSTEM \"%s\">\n", statics.DTDFile)

	members := make([]Member, len(users))
	for i := range users {
		user := users[i]
		members[i] = Member{
			ID:   user.Arma.ID,
			Name: user.Name,
			Nick: user.Arma.Nick,
		}
	}

	squad := Squad{
		Nick:    "Anrop",
		Name:    "Anrop",
		Email:   "admin@anrop.se",
		Picture: statics.PAAFile,
		Title:   "Anrop",
		Web:     "https://www.anrop.se",
		Members: members,
	}

	output, err := xml.Marshal(squad)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling XML: %q\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write([]byte(xml.Header + doctype + string(output)))
}
