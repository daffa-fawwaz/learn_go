package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type inputUser struct {
	kota     string
	suhu     float32
	kategori string
}

func convertTemperature(user *inputUser) {
	reamur := (user.suhu * 4) / 5
	fahrenheit := (user.suhu * 1.8) + 32

	fmt.Printf("Suhu di %s adalah: %s\n", user.kota, user.kategori)
	fmt.Printf("Suhu dalam Reamur: %.2f\n", reamur)
	fmt.Printf("Suhu dalam Fahrenheit: %.2f\n", fahrenheit)
}

func checkCondition(celcius float32) string {
	if celcius < 18 {
		return "dingin"
	} else if celcius > 25 {
		return "panas"
	} else {
		return "normal"
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan kota anda: ")
	kota, _ := reader.ReadString('\n')
	kota = strings.TrimSpace(kota)
	_, errKota := strconv.ParseFloat(kota, 32)
	if errKota == nil {
		fmt.Println("Masukkan nama kota, bukan angka.")
		return
	}

	fmt.Print("Masukkan suhu dalam Celcius: ")
	suhu, _ := reader.ReadString('\n')
	suhu = strings.TrimSpace(suhu)
	celciusFloat, errSuhu := strconv.ParseFloat(suhu, 32)
	if errSuhu != nil {
		fmt.Println("Input tidak valid. Harus berupa angka.")
		return
	}

	user := inputUser{
		kota:     kota,
		suhu:     float32(celciusFloat),
		kategori: checkCondition(float32(celciusFloat)),
	}

	convertTemperature(&user)

}
