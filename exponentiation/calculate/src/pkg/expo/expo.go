package expo

import (
	"log"
	"math"
)

type Inputs struct {
	Base float64
	Exponent float64
}

func (in Inputs) Run() {
	calcPow := math.Pow(in.Base, in.Exponent)
	log.Println("Starting the calculation...")
	log.Printf("%f ^ %f\n", in.Base, in.Exponent)
	log.Printf("Result of exponetiation: %f\n", calcPow)
	log.Println("Finished calculation.")
}