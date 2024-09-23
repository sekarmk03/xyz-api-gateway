package utils

import (
	"log"
	"strconv"
)

func StrParamToInt(s string, def int) int {
	result, err := strconv.Atoi(s)
	if err != nil || result < 1 {
		log.Printf("ERROR: [utils - StrParamToInt] Error while convert %s string to int: %v\n", s, err)
		result = def
	}
	return result
}
