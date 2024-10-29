// algorithm implements an image analysis algorithm that complies with the MdTF image analysis API.
// Its analysis output is a pseudorandom function of partial image contents.
package algorithm

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"

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

	results.Score, results.NormalizedScore = genRandomImageScore(imageBytes)
}

// genRandomImageScore generates a random example analysis score.
// Output is seeded by image contents to maintain deterministic stateless algorithm behavior.
func genRandomImageScore(imageBytes []byte) (score string, normalizedScore string) {
	var imagePrefix []byte

	if len(imageBytes) >= 64 {
		imagePrefix = imageBytes[len(imageBytes)-64:]
	} else {
		imagePrefix = imageBytes
	}
	imageSeed, _ := binary.Varint(imagePrefix)
	rand.Seed(imageSeed)
	scoreFloat := math.Min(rand.ExpFloat64() / 0.5, 10.0)
	normalizedScoreFloat := scoreFloat / 10.0
	return fmt.Sprintf("%.3f", scoreFloat), fmt.Sprintf("%.3f", normalizedScoreFloat)
}

// truncateOutput ensures a basic response will be returned if the response exceeds its maximum allowable size.
func truncateOutput(ar models.AnalysisResult) (models.AnalysisResult, error){
	responseBytes, err := json.Marshal(ar)
	if err != nil{
		return models.AnalysisResult{}, errors.New("error preparing response: " + err.Error())
	}

	if len(responseBytes) > 512{
		return models.AnalysisResult{ Score: ar.Score, NormalizedScore: ar.NormalizedScore, AnalysisError: "truncated" }, nil
	}
	return ar, nil
}
