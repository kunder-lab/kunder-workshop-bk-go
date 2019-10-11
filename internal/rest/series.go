package rest

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Series struct {
	Series []Serie `json:"series"`
}
type Serie struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ID          string   `json:"id"`
	Images      []string `json:"images"`
	Type        string   `json:"type"`
	NumEpisodes int      `json:"numEpisodes"`
}

// home is a simple HTTP handler function which writes a response.
func getSeries(w http.ResponseWriter, _ *http.Request) {
	requestURL := "http://demo7835208.mockable.io/series"

	resp, _ := http.Get(requestURL)

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var series Series
	err := json.Unmarshal(bodyBytes, &series)
	if err != nil {
		log.Fatal("Series no trae formato correcto", err)
	}
	log.Println(series.Series[0].Description)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bodyBytes)
	w.WriteHeader(http.StatusOK)
}
