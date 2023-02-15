package models

import "fmt"

type Gls struct {
	QrCode string
}

func (gl Gls) record() {
	fmt.Println("ADD A GLS RECORD")
}
