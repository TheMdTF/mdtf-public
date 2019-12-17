// algorithm implements an image analysis algorithm that complies with the MdTF image analysis API.
// Its numeric analysis score output is a pseudorandom function of partial image contents.
package algorithm

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/TheMdTF/mdtf-public/image-analysis/go-example/models"
)

// AnalyzeImage runs image analysis and returns the full JSON output for a single input image.
func AnalyzeImage(imageBytes[]byte) (models.AnalysisResult, error) {
	resultsContainer := models.AnalysisResult{}
	analyzeHelper(imageBytes, &resultsContainer)
	return truncateOutput(resultsContainer)
}

func analyzeHelper(imageBytes []byte, results *models.AnalysisResult) {
	// recover from any panic and store the message as AnalysisError field of output
	defer func(ar *models.AnalysisResult){
		if r := recover(); r != nil{
				analyzeError, ok := r.(error)
				if ok{
					ar.AnalysisError = analyzeError.Error()
				} else{
					ar.AnalysisError = "unknown error"
				}
			}
		}(results)

	start := time.Now()
	results.Score = genRandomImageScore(imageBytes)
	compTime := time.Since(start)

	results.Timestamp = start.String()

	compTimeMs := float32(compTime.Nanoseconds()) / float32(time.Millisecond.Nanoseconds())
	results.ComputationTimeMilli = fmt.Sprintf("%f", compTimeMs)
}

// genRandomImageScore generates a random example AnalysisScore.
// Output is seeded by image contents to maintain deterministic stateless algorithm behavior.
func genRandomImageScore(imageBytes []byte) float32 {
	var imagePrefix []byte

	// Sleep to simulate algorithm computation time
	rand.Seed(time.Now().Unix() - int64(len(imageBytes)))
	time.Sleep(time.Millisecond * time.Duration(rand.Float32()*25))

	if len(imageBytes) >= 64 {
		imagePrefix = imageBytes[len(imageBytes)-64:]
	} else {
		imagePrefix = imageBytes
	}
	imageSeed, _ := binary.Varint(imagePrefix)
	rand.Seed(imageSeed)
	return rand.Float32()
}

// truncateOutput ensures a basic response will be returned if the response exceeds its maximum allowable size.
func truncateOutput(ar models.AnalysisResult) (models.AnalysisResult, error){
	responseBytes, err := json.Marshal(ar)
	if err != nil{
		return models.AnalysisResult{}, errors.New("error preparing response: " + err.Error())
	}

	if len(responseBytes) > 512{
		return models.AnalysisResult{ Score: ar.Score, Timestamp: ar.Timestamp, AnalysisError: "truncated" }, nil
	}
	return ar, nil
}
