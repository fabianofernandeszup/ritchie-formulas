package main

import (
	"expo/pkg/expo"
	"os"
	"strconv"
)

func main() {
	base := os.Getenv("BASE")
	exponent := os.Getenv ("EXPONENT")
	baseFloat, _ := strconv.ParseFloat(base, 64)
	exponentFloat, _ := strconv.ParseFloat(exponent,64)
	expo.Inputs{
		Base: baseFloat,
		Exponent: exponentFloat,
	}.Run()
}