package vars

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MandatoryString(variableName string) (string, error) {
	value := os.Getenv(variableName)
	if value == "" {
		return "", fmt.Errorf("unable to find variable: %s ", variableName)
	}
	return value, nil
}

func MandatoryInt(variableName string) (int, error) {
	value := os.Getenv(variableName)
	if value == "" {
		return 0, fmt.Errorf("unable to find variable: %s ", variableName)
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func MandatoryBool(variableName string) (bool, error) {
	value := os.Getenv(variableName)
	if value == "" {
		return false, fmt.Errorf("unable to find variable: %s ", variableName)
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}
	return b, nil
}

func MandatoryFloat32(variableName string) (float32, error) {
	value := os.Getenv(variableName)
	if value == "" {
		return 0, fmt.Errorf("unable to find variable: %s ", variableName)
	}
	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

func MandatoryFloat64(variableName string) (float64, error) {
	value := os.Getenv(variableName)
	if value == "" {
		return 0, fmt.Errorf("unable to find variable: %s ", variableName)
	}
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

func MandatoryArrayString(variableName string) ([]string, error) {
	value := os.Getenv(variableName)
	if value == "" {
		return nil, fmt.Errorf("unable to find variable: %s ", variableName)
	}
	// split the value by comma
	arrVal := strings.Split(value, ",")
	return arrVal, nil
}

func MandatoryArrayInt(variableName string) ([]int, error) {
	value := os.Getenv(variableName)
	if value == "" {
		return nil, fmt.Errorf("unable to find variable: %s ", variableName)
	}
	// split the value by comma
	arrVal := strings.Split(value, ",")
	var arrInt []int
	for _, val := range arrVal {
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		arrInt = append(arrInt, i)
	}
	return arrInt, nil
}
