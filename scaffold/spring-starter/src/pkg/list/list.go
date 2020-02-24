package list

import (
	"fmt"
	"log"
)

type Inputs struct {
	Command string
}

func (in Inputs) Run() {
	fmt.Printf("Command: %v\n", in.Command)
	log.Println("Listar dependências")
}
