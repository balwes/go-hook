package util

import (
	"strconv"
)

func PanicIfNotNil(e error) {
	if e != nil {
		panic(e)
	}
}

func PanicIfFalse(b bool) {
	if !b {
		panic("Expected a true bool")
	}
}

func FloatToString(f float32) string {
    return strconv.FormatFloat(float64(f), 'f', -1, 32)
}
