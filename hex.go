package main

import "strconv"

func hexToDecimal(hex string) (int) {
	number, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		
	}
	return int(number)
}