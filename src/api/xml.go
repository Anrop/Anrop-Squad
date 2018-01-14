package api

import (
	"db"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
)

var doctype = "<!DOCTYPE squad SYSTEM \"squad.dtd\">\n"

func Arma2XmlHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetArma2Users()
	internalXmlHandler(w, r, *users, err)
}

func Arma3XmlHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetArma3Users()
	internalXmlHandler(w, r, *users, err)
}

func internalXmlHandler(w http.ResponseWriter, r *http.Request, users []db.User, err error) {
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

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
		Picture: "squad.paa",
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
