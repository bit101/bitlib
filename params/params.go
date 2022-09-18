package params

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type Param struct {
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Value float64 `json:"value"`
}

type Params map[string]*Param

var params Params

func addParam(name string, min, max, value float64) {
	params[name] = &Param{min, max, value}
}

func LoadParams(filepath string) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(bytes, &params)
}

func GetValue(name string) float64 {
	param, ok := params[name]
	if !ok {
		log.Fatalf("No param named %s\n", name)
	}
	return param.Value
}

func GetMin(name string) float64 {
	param, ok := params[name]
	if !ok {
		log.Fatalf("No param named %s\n", name)
	}
	return param.Min
}

func GetMax(name string) float64 {
	param, ok := params[name]
	if !ok {
		log.Fatalf("No param named %s\n", name)
	}
	return param.Max
}

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

func List() []string {
	list := make([]string, len(params))
	i := 0
	for k := range params {
		list[i] = k
		i++
	}
	return list
}
