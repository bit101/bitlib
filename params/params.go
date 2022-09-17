package params

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Param struct {
	min   float64
	max   float64
	value float64
}

var params = make(map[string]*Param)

func addParam(name string, min, max, value float64) {
	params[name] = &Param{min, max, value}
}

func GetValue(name string) float64 {
	param, ok := params[name]
	if !ok {
		log.Fatalf("No param named %s\n", name)
	}
	return param.value
}

func GetMin(name string) float64 {
	param, ok := params[name]
	if !ok {
		log.Fatalf("No param named %s\n", name)
	}
	return param.min
}

func GetMax(name string) float64 {
	param, ok := params[name]
	if !ok {
		log.Fatalf("No param named %s\n", name)
	}
	return param.max
}

func GetParam(name string) (float64, float64, float64) {
	param, ok := params[name]
	if !ok {
		log.Fatalf("No param named %s\n", name)
	}
	return param.min, param.max, param.value
}

func LoadParams(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Split(line, " ")
		if len(parts) == 4 {
			addParam(parts[0], getFloat(parts[1]), getFloat(parts[2]), getFloat(parts[3]))
		}
	}
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
