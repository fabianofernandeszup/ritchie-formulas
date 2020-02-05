package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"topic/pkg/topic"
)

func main() {
	i, err := loadInputs()
	if err != nil {
		fmt.Println(err)
		return
	}

	topic.Create(i)
}

func loadInputs() (*topic.Inputs, error) {
	u := os.Getenv("URLS")
	n := os.Getenv("NAME")
	r := os.Getenv("REPLICATION")
	p := os.Getenv("PARTITIONS")

	re, err := strconv.Atoi(r)
	if err != nil {
		return nil, errors.New("replication must be a number")
	}

	pa, err := strconv.Atoi(p)
	if err != nil {
		return nil, errors.New("partitions must be a number")
	}

	i := topic.Inputs{Urls: u, Name: n, Replication: int16(re), Partitions: int32(pa)}

	return &i, nil
}
