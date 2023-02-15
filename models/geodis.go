package models

import "fmt"

type Geodis struct {
	PrintingLabels string
}

func (geo Geodis) record() {
	fmt.Println("ADD A GEODIS RECORD")
}
