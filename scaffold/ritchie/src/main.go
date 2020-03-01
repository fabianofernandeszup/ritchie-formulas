package main

import (
	"os"
	"ritchie/pkg/ritchie"
)

func main() {

	 input := os.Getenv("EXAMPLE_ENTRY")

	ritchie.Send(ritchie.Msg{Value: "Hello World."})
	ritchie.Send(ritchie.Msg{Value: input})
}
