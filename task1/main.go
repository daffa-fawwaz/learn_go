package main

import "fmt"

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
	var celcius float32
	fmt.Print("Masukkan suhu dalam celcius: ")
	_, err := fmt.Scan(&celcius)

	if err != nil {
		fmt.Println("Input tidak valid. Harus berupa angka.")
		return
	}

	if err == nil {
		fmt.Println("Suhu dalam reamur: ", celciusToReamur(celcius))
		fmt.Println("Suhu dalam fahrenheit: ", celciusToFahrenheit(celcius))
	}

}
