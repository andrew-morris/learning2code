package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func whatTheFuck() {
	type Dog struct {
		Breed  string
		Name   string
		Weight int
		Color  string
	}
	Vinnie := Dog{
		Breed:  "pitbull",
		Name:   "vinnie",
		Weight: 60,
		Color:  "tan",
	}
	b, err := json.Marshal(Vinnie)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(b)
}

func main() {
	whatTheFuck()
}
