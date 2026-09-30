package main

import "strconv"

//converts hexadecimals to decimals
func hexToDecimal(hex string) (int, error) {
	number, err := strconv.ParseInt(hex, 16, 64)
	return int(number), err
}

//converts binary to decimals
func binToDecimal(bin string) (int, error) {
	number, err := strconv.ParseInt(bin, 2, 64)
    return int(number), err
}