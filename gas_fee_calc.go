package main

import "fmt"

func calcGasFee(gasLimit int, gasPrice float64) float64 {
	return float64(gasLimit) * gasPrice
}

func main() {
	fee := calcGasFee(21000, 0.00001)
	fmt.Printf("燃料费：%.6f\n", fee)
}
