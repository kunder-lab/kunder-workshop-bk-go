package rest

import (
	"fmt"
	"kunder-workshop-bk-go/settings"
	"net/http"
)

// home is a simple HTTP handler function which writes a response.
func home(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Version "+settings.V.Version)
}
