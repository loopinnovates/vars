package vars

import (
	"os"
	"strconv"
	"strings"
)

func OptionalString(variableName string, defaultValue string) string {
	value := os.Getenv(variableName)
	if value == "" {
		return defaultValue
	}
	return value
}

func OptionalInt(variableName string, defaultValue int) int {
	value := os.Getenv(variableName)
	if value == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i
}

func OptionalBool(variableName string, defaultValue bool) bool {
	value := os.Getenv(variableName)
	if value == "" {
		return defaultValue
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return b
}

func OptionalFloat32(variableName string, defaultValue float32) float32 {
	value := os.Getenv(variableName)
	if value == "" {
		return defaultValue
	}
	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return defaultValue
	}
	return float32(f)
}

func OptionalFloat64(variableName string, defaultValue float64) float64 {
	value := os.Getenv(variableName)
	if value == "" {
		return defaultValue
	}
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultValue
	}
	return f
}

func OptionalArrayString(variableName string, defaultValue []string) []string {
	value := os.Getenv(variableName)
	if value == "" {
		return defaultValue
	}
	// split the value by comma
	arrVal := strings.Split(value, ",")
	return arrVal
}

func OptionalArrayInt(variableName string, defaultValue []int) []int {
	value := os.Getenv(variableName)
	if value == "" {
		return defaultValue
	}
	// split the value by comma
	arrVal := strings.Split(value, ",")
	var arrInt []int
	for _, val := range arrVal {
		i, err := strconv.Atoi(val)
		if err != nil {
			return defaultValue
		}
		arrInt = append(arrInt, i)
	}
	return arrInt
}

func OptionalArrayFloat64(variableName string, defaultValue []float64) []float64 {
	value := os.Getenv(variableName)
	if value == "" {
		return defaultValue
	}
	// split the value by comma
	arrVal := strings.Split(value, ",")
	var arrFloat []float64
	for _, val := range arrVal {
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return defaultValue
		}
		arrFloat = append(arrFloat, f)
	}
	return arrFloat
}

func OptionalArrayAny(variableName string, defaultValue any) any {
	value := os.Getenv(variableName)
	if value == "" {
		return defaultValue
	}
	// split the value by comma
	arrVal := strings.Split(value, ",")
	var arrAny []any
	for _, val := range arrVal {
		arrAny = append(arrAny, val)
	}
	return arrAny
}
