// Copyright 2018 The Maryland Test Facility

package main

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -L./lib -lsample_matching_algorithm
// #include <stdlib.h>
// #include "./c/sample_matching_algorithm.h"
import "C"

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unsafe"

	"github.com/TheMdTF/mdtf-public/rally2-matching-system/go-example/models"
	"github.com/disintegration/imaging"
)

func createTemplate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		log.Println("received call to create template")

		//parse the json
		jsonDecoder := json.NewDecoder(r.Body)
		var imageModel models.Image
		err := jsonDecoder.Decode(&imageModel)
		if err != nil {
			log.Println("error decoding create template request: ", err)
			http.Error(w, "Bad Image Model: "+err.Error(), http.StatusBadRequest)
			return
		}

		//decode the image string
		imageByteData, err := base64.StdEncoding.DecodeString(imageModel.ImageData)
		if err != nil {
			log.Println("error decoding base64 image string to bytes: ", err)
			http.Error(w, "Bad Base 64 Encoding: "+err.Error(), http.StatusBadRequest)
			return
		}

		//check for valid image data
		reader := bytes.NewReader(imageByteData)
		_, err = imaging.Decode(reader)
		if err != nil {
			log.Println("error decoding bytes as image data: ", err)
			http.Error(w, "Bad Image Data: "+err.Error(), http.StatusBadRequest)
			return
		}

		//check that the valid image is a png
		if !strings.HasPrefix(fmt.Sprintf("%s", imageByteData), "\x89PNG\r\n\x1a\n") {
			log.Println("invalid image type (not a PNG): ", err)
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
		err = json.NewEncoder(w).Encode(template)
		if err != nil {
			log.Println("error encoding response: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func compareList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		log.Println("received call to compare list")

		//parse the json
		decoder := json.NewDecoder(r.Body)
		var compRequest models.CompareListRequest
		err := decoder.Decode(&compRequest)
		if err != nil {
			log.Println("error decoding compare list request: ", err)
			http.Error(w, "Bad Comparison Request Data", http.StatusBadRequest)
			return
		}
		log.Printf("list contains %d templates\n", len(compRequest.TemplateList))

		//pass each comparison to the C library
		var cList []models.Comparison
		var decoded1 []byte
		decoded1, err = base64.StdEncoding.DecodeString(compRequest.SingleTemplate.Template)
		if err != nil {
			log.Println("error decoding single template: ", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		template1 := C.CString(string(decoded1))
		defer C.free(unsafe.Pointer(template1))
		for i := 0; i < len(compRequest.TemplateList); i++ {
			var decoded2 []byte
			decoded2, err = base64.StdEncoding.DecodeString(compRequest.TemplateList[i].Template)
			if err != nil {
				log.Printf("error decoding the template at index %d: %s", i, err.Error())
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			template2 := C.CString(string(decoded2))
			defer C.free(unsafe.Pointer(template2))

			s := C.cpp_compare_template(template1, template2)

			cList = append(cList, models.Comparison{
				Score:           float32(s),
			})
		}

		//return response
		err = json.NewEncoder(w).Encode(cList)
		if err != nil {
			log.Println("error encoding response: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func info(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("received call to info")

		i := models.Info{
			AlgorithmName:         "Example MdTF Matching Algorithm",
			AlgorithmVersion:      "1.0.0",
			AlgorithmType:         "Face",
			CompanyName:           "MdTF",
			TechnicalContactEmail: "john@mdtf.org",
			RecommendedCPUs:       4,	
			RecommendedMem:        2048,
			Test:                  "MDTF_2021_RALLY",
			Thresholds: map[string]float32{
				"1:500": float32(0.75),
				"1:1e3": float32(0.85),
				"1:1e4": float32(0.95),
				"1:1e5": float32(0.97),
				"1:1e6": float32(0.99),
			},
		}

		err := json.NewEncoder(w).Encode(i)
		if err != nil {
			log.Println("error encoding response: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

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
