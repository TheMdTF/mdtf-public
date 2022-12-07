package models

import (
	"encoding/json"
	"fmt"
	"errors"
	"strconv"
	"strings"
)

// Info Basic information describing the algorithm.
type Info struct {

	// Name of algorithm
	AlgorithmName string

	// A string enum describing the type of biometric images the algorithm is meant to process
	// Enum: [Face Finger Iris]
	AlgorithmType string

	// Algorithm version identifier
	AlgorithmVersion string

	// Name of the Company which produces the algorithm
	CompanyName string

	// The recommended allocation of CPUs for the deployed docker container.
	RecommendedCPUs int64

	// The recommended allocation of memory (MB) for the deployed docker container.
	RecommendedMem int64

	// The email address of an engineer or other technical resource to contact in the event of an error running your service.
	// This field may be left blank if desired.
	TechnicalContactEmail string

	// A string enum describing which test to run with the algorithm
	Test string

	// The list of threshold scores for various false match rates (FMRs) for your submitted matching system.
	// Systems are required to provide at least thresholds for achieving FMRs for 1:500, 1:1e3,  1:1e4,  1:1e5,  1:1e6
	Thresholds map[string]float32
}

//MarshalJSON implements the Marshaller interface to meet swagger requirements
func (i Info) MarshalJSON() (b []byte, err error) {
	temp := struct {
		AlgorithmName         string
		AlgorithmType         string
		AlgorithmVersion      string
		CompanyName           string
		RecommendedCPUs       int64
		RecommendedMem        int64
		TechnicalContactEmail string
		Test                  string
		Thresholds            map[string]string
	}{
		AlgorithmName:         i.AlgorithmName,
		AlgorithmType:         i.AlgorithmType,
		AlgorithmVersion:      i.AlgorithmVersion,
		CompanyName:           i.CompanyName,
		RecommendedCPUs:       i.RecommendedCPUs,
		RecommendedMem:        i.RecommendedMem,
		TechnicalContactEmail: i.TechnicalContactEmail,
		Test:                  i.Test,
		Thresholds:            convertMapString(i.Thresholds),
	}
	return json.Marshal(temp)
}

func convertMapString(in map[string]float32) (out map[string]string) {
	out = make(map[string]string, len(in))
	for k, v := range in {
		out[k] = strings.TrimSpace(fmt.Sprintf("%10f", v))
	}
	return
}

//UnmarshalJSON implements the Unmarshaller interface to meet swagger requirements
func (i *Info) UnmarshalJSON(b []byte) (err error) {
	temp := struct {
		AlgorithmName         string
		AlgorithmType         string
		AlgorithmVersion      string
		CompanyName           string
		RecommendedCPUs       int64
		RecommendedMem        int64
		TechnicalContactEmail string
		Test                  string
		Thresholds            map[string]string
	}{}
	err = json.Unmarshal(b, &temp)
	if err != nil {
		return
	}

	thresh, err := convertMapFloat(temp.Thresholds)
	if err != nil {
		return
	}
	if len(thresh) > 10 {
		return errors.New("too many thresholds given, should be at most 10")
	}

	i.AlgorithmName = temp.AlgorithmName
	i.AlgorithmType = temp.AlgorithmType
	i.AlgorithmVersion = temp.AlgorithmVersion
	i.CompanyName = temp.CompanyName
	i.RecommendedCPUs = temp.RecommendedCPUs
	i.RecommendedMem = temp.RecommendedMem
	i.TechnicalContactEmail = temp.TechnicalContactEmail
	i.Test = temp.Test
	i.Thresholds = thresh
	return
}

func convertMapFloat(in map[string]string) (out map[string]float32, err error) {
	out = make(map[string]float32, len(in))
	for k, v := range in {
		if len(v) > 10 {
			return nil, fmt.Errorf("invalid threshold length; was: %d and should be at max 10", len(v))
		}
		var val float64
		val, err = strconv.ParseFloat(v, 32)
		if err != nil {
			return
		}
		out[k] = float32(val)
	}
	return
}
