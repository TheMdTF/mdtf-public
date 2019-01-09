package main

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -L./lib -lsample_matching_algorithm
// #include <stdlib.h>
// #include "./c/sample_matching_algorithm.h"
import "C"

import (
	"net/http"
	"encoding/json"
	"unsafe"
	"github.com/TheMdTF/mdtf-public/rally2-matching-system/go-example/models"
	"encoding/base64"
	"strconv"
	"github.com/disintegration/imaging"
	"bytes"
	"strings"
	"fmt"
)

func createTemplate(w http.ResponseWriter, r *http.Request) {

	switch r.Method{
	case http.MethodPost:

		//parse the json
		jsonDecoder := json.NewDecoder(r.Body)
		var imageModel models.Image
		err := jsonDecoder.Decode(&imageModel)
		if err != nil {
			http.Error(w, "Bad Image Model: " + err.Error(), http.StatusBadRequest)
			return
		}

		//decode the image string
		imageByteData, err := base64.StdEncoding.DecodeString(imageModel.ImageData)
		if err != nil {
			http.Error(w, "Bad Base 64 Encoding: " + err.Error(), http.StatusBadRequest)
			return
		}

		//check for valid image data
		reader := bytes.NewReader(imageByteData)
		_, err = imaging.Decode(reader)
		if err != nil {
			http.Error(w, "Bad Image Data: " + err.Error(), http.StatusBadRequest)
			return
		}

		//check that the valid image is a png
		if !strings.HasPrefix(fmt.Sprintf("%s", imageByteData), "\x89PNG\r\n\x1a\n"){
			http.Error(w, "Image Data not a PNG", http.StatusBadRequest)
			return
		}

		//pass image to C library
		imageCData := C.CString(imageModel.ImageData)
		defer C.free(unsafe.Pointer(imageCData))

		r := C.cpp_create_template(imageCData)

		b := strconv.FormatInt(int64(r), 2)
		template := models.Template{
			Template: base64.StdEncoding.EncodeToString([]byte(b)),
		}

		//return response
		json.NewEncoder(w).Encode(template)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func compareList(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case http.MethodPost:

		//parse the json
		decoder := json.NewDecoder(r.Body)
		var compRequest models.CompareListRequest
		err := decoder.Decode(&compRequest)
		if err != nil {
			http.Error(w, "Bad Comparison Request Data", http.StatusBadRequest)
			return
		}

		//pass each comparison to the C library
		var cList [] models.Comparison
		decoded1, _ := base64.StdEncoding.DecodeString(compRequest.SingleTemplate.Template)
		template1 := C.CString(string(decoded1))
		defer C.free(unsafe.Pointer(template1))
		for i := 0; i < len(compRequest.TemplateList); i++ {
			decoded2, _ := base64.StdEncoding.DecodeString(compRequest.TemplateList[i].Template)
			template2 := C.CString(string(decoded2))
			defer C.free(unsafe.Pointer(template2))

			s := C.cpp_compare_template(template1, template2)

			cList = append(cList, models.Comparison{
				Score: float32(s),
				NormalizedScore: float32(s),
			})
		}

		//return response
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
