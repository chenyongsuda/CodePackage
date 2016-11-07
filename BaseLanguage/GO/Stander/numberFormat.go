package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
Format int to string with 0 prefix
sample 888 -> 00000000888
*/

var once sync.Once

type NumberFormat struct {
	digiestWidth int
}

var instanceNumberFormat *NumberFormat

func NewNumberFormat() *NumberFormat {
	return &NumberFormat{digiestWidth: 20}
}

func GetNumberFormatInstance() *NumberFormat {
	once.Do(func() {
		instanceNumberFormat = &NumberFormat{digiestWidth: 20}
	})
	return instanceNumberFormat
}

func (nf *NumberFormat) format(v int) string {
	fmtStr := "%0" + strconv.Itoa(nf.digiestWidth) + "d"
	return fmt.Sprintf(fmtStr, v)
}

func main() {
	// test := NewNumberFormat()

	test := GetNumberFormatInstance()
	fmt.Println(test.format(888))
	test2 := GetNumberFormatInstance()
	fmt.Println(test2.format(666))
}
