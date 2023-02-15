package models

import "fmt"

type Dhl struct {
	Nfc string
}

func (dh Dhl) record() {
	fmt.Println("ADD A DHL RECORD")
}
