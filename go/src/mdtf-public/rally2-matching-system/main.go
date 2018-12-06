package main

import (
	"net/http"
	"encoding/json"
	"mdtf-public/rally2-matching-system/models"
)

func createTemplate(w http.ResponseWriter, r *http.Request) {

	switch r.Method{
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		var img models.Image
		err := decoder.Decode(&img)
		if err != nil {
			http.Error(w, "Bad Image Data", http.StatusBadRequest)
		}

		w.Write([]byte("Creating Template " + img.ImageData + "\n"))

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func compareList(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case http.MethodPost:
		w.Write([]byte("Comparing List\n"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/v1/create-template", createTemplate)
	http.HandleFunc("/v1/compare-list", compareList)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}


