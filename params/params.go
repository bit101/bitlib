// Package params reads params from a json file.
package params

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

// Param represents a single parameter.
type Param struct {
	Min      float64 `json:"min"`
	Max      float64 `json:"max"`
	Value    float64 `json:"value"`
	Decimals int     `json:"decimals"`
}

// Params is a list of Params.
type Params map[string]*Param

var params Params

// LoadParams loads the params from an external json file.
func LoadParams(filepath string) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(bytes, &params)
}

// GetValue gets a named param's value.
func GetValue(name string) float64 {
	param, ok := params[name]
	if !ok {
		log.Fatalf("No param named %s\n", name)
	}
	return param.Value
}

// GetMin gets the minimum value of a named param.
func GetMin(name string) float64 {
	param, ok := params[name]
	if !ok {
		log.Fatalf("No param named %s\n", name)
	}
	return param.Min
}

// GetMax gets the maximum value of a given param.
func GetMax(name string) float64 {
	param, ok := params[name]
	if !ok {
		log.Fatalf("No param named %s\n", name)
	}
	return param.Max
}

// GetDecimals gets the number of decimals for a given parameter.
func GetDecimals(name string) int {
	param, ok := params[name]
	if !ok {
		log.Fatalf("No param named %s\n", name)
	}
	return param.Decimals
}

// GetParam gets the min, max and value of a given parameter.
func GetParam(name string) (float64, float64, float64) {
	param, ok := params[name]
	if !ok {
		log.Fatalf("No param named %s\n", name)
	}
	return param.Min, param.Max, param.Value
}

func getFloat(str string) float64 {
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatalf("Can't parse '%s' as a float64", str)
	}
	return value
}

// List lists all parameters.
func List() []string {
	list := make([]string, len(params))
	i := 0
	for k := range params {
		list[i] = k
		i++
	}
	return list
}
