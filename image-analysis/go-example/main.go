package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/TheMdTF/mdtf-public/image-analysis/go-example/algorithm"
	"github.com/TheMdTF/mdtf-public/image-analysis/go-example/models"
	"github.com/disintegration/imaging"
)

func info(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("received call to info")

		i := models.Info{
			AlgorithmName:         "MdTF Image Analysis Algorithm",
			AlgorithmType:         "Face",
			AlgorithmVersion:      "1.0.0",
			CompanyName:           "MdTF",
			ImageDataset:		   "Enrollment",
			RecommendedCPUs:       4,
			RecommendedMem:        2048,
			TechnicalContactEmail: "john@mdtf.org",
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

func analyzeImage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		log.Println("received call to analyze-image")

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

		//check that the valid image is a png
		if !strings.HasPrefix(fmt.Sprintf("%s", imageByteData), "\x89PNG\r\n\x1a\n") {
			log.Println("invalid image type (not a PNG): ", err)
			http.Error(w, "Image Data not a PNG", http.StatusBadRequest)
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

		//run analysis algorithm on image
		analysisResult, err := algorithm.AnalyzeImage(imageByteData)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//return response
		err = json.NewEncoder(w).Encode(analysisResult)
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
	http.HandleFunc("/v1/info", info)
	http.HandleFunc("/v1/analyze-image", analyzeImage)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
