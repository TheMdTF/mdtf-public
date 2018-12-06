package main

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -L../../../../lib -lsample_matching_algorithm
// #include <stdlib.h>
// #include "../../../../c/sample_matching_algorithm.h"
import "C"

import (
	"net/http"
	"encoding/json"
	"unsafe"
	"mdtf-public/rally2-matching-system/models"
	b64 "encoding/base64"
	"strconv"
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

		image := C.CString(img.ImageData)
		defer C.free(unsafe.Pointer(image))

		r := C.cpp_create_template(image)

		b := strconv.FormatInt(int64(r), 2)

		template := models.Template{
			Template: b64.StdEncoding.EncodeToString([]byte(b)),
		}

		json.NewEncoder(w).Encode(template)

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


