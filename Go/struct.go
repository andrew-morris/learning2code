package main

import (
	"fmt"
)

type Band struct {
	singer      string
	guitarist   string
	bassist     string
	keyboardist string
	drummer     string
}

func main() {
	MyLosingSeason := Band{singer: "colin",
		guitarist:   "josh",
		bassist:     "david",
		keyboardist: "coleman",
		drummer:     "andrew",
	}

	fmt.Println(MyLosingSeason.guitarist)
}
