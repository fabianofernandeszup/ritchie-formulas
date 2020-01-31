package main

import (
	"os"
	"strconv"

	"consume/pkg/consume"
)

func main() {
	consume.Consume(loadInputs())
}

func loadInputs() consume.Inputs {
	b, _ := strconv.ParseBool(os.Getenv("BEGINNING"))

	return consume.Inputs{
		Url:           os.Getenv("URL"),
		Topic:         os.Getenv("TOPIC"),
		FromBeginning: b,
	}
}
