package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func celciusToReamur(celcius float32) float32 {
	var reamur float32
	reamur = (celcius * 4) / 5
	return reamur
}

func celciusToFahrenheit(celcius float32) float32 {
	var fahrenheit float32
	fahrenheit = (celcius * 1.8) + 32
	return fahrenheit
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan suhu dalam Celcius: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	celciusFloat, err := strconv.ParseFloat(input, 32)

	if err != nil {
		fmt.Println("Input tidak valid. Harus berupa angka.")
		return
	}

	celcius := float32(celciusFloat)

	fmt.Println("Suhu dalam reamur: ", celciusToReamur(celcius))
	fmt.Println("Suhu dalam fahrenheit: ", celciusToFahrenheit(celcius))

}
