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
		decoder := json.NewDecoder(r.Body)
		var compRequest models.CompareListRequest
		err := decoder.Decode(&compRequest)
		if err != nil {
			http.Error(w, "Bad Comparison Request Data", http.StatusBadRequest)
		}

		var cList [] models.Comparison
		template1 := C.CString(compRequest.SingleTemplate.Template)
		defer C.free(unsafe.Pointer(template1))
		for i := 0; i < len(compRequest.TemplateList); i++ {
			template2 := C.CString(compRequest.TemplateList[i].Template)
			defer C.free(unsafe.Pointer(template2))

			s := C.cpp_compare_template(template1, template2)

			cList = append(cList, models.Comparison{
				Score: float32(s),
				NormalizedScore: float32(s),
			})
		}

		json.NewEncoder(w).Encode(cList)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func info(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case http.MethodGet:
		i := models.Info{
			AlgorithmName: "Example MdTF Matching Algorithm",
			AlgorithmVersion: "1.0.0",
			AlgorithmType: "Iris",
			CompanyName: "MdTF",
			TechnicalContactEmail: "john@mdtf.org",
			RecommendedCpus: 4,
			RecommendedMem: 2048,
		}

		json.NewEncoder(w).Encode(i)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/v1/create-template", createTemplate)
	http.HandleFunc("/v1/compare-list", compareList)
	http.HandleFunc("/v1/info", info)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
