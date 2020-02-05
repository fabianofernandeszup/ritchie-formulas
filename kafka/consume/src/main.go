package main

import (
	"os"
	"strconv"

	"consume/pkg/consume"
)

func main() {
	consume.Run(loadInputs())
}

func loadInputs() consume.Inputs {
	b, _ := strconv.ParseBool(os.Getenv("BEGINNING"))
	return consume.Inputs{
		Urls:          os.Getenv("URLS"),
		Topic:         os.Getenv("TOPIC"),
		FromBeginning: b,
	}
}
