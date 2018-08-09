package main

import (
	"fmt"
	"strconv"
)

func float32To(f float32) string {
	if i, err := strconv.Atoi(fmt.Sprintf("%0.0f", f)); err == nil && f-float32(i) == 0 {
		return "int"
	}
	return "float32"
}

func float64To(f float64) string {
	if i, err := strconv.Atoi(fmt.Sprintf("%0.0f", f)); err == nil && f-float64(i) == 0 {
		return "int"
	}
	return "float64"
}
