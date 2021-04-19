package api

import (
	"github.com/Anrop/Anrop-Squad/db"
	"github.com/Anrop/Anrop-Squad/static"
	"io"
	"net/http/httptest"
	"testing"
)

var statics = static.Statics{
	DTDFile: "squad-123.dtd",
	FS:      nil,
	PAAFile: "squad-123.paa",
}

var users = []db.User{
	{
		ID:   "User ID",
		Name: "User Name",
		Arma: db.Arma{
			ID:   "Arma ID",
			Nick: "Arma Nick",
		},
	},
}

func TestXMLHandler(t *testing.T) {
	expected := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE squad SYSTEM "squad-123.dtd">
<squad nick="Anrop"><email>admin@anrop.se</email><name>Anrop</name><picture>squad-123.paa</picture><title>Anrop</title><web>https://www.anrop.se</web><member id="Arma ID" nick="Arma Nick"><name>User Name</name></member></squad>`

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/arma3.xml", nil)

	internalXMLHandler(w, r, &statics, users, nil)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	bodyStr := string(body)

	if bodyStr != expected {
		t.Fatalf("Expected did not match body: %s", body)
	}
}
