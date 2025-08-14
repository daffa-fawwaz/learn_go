package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FormPendaftaranInterface interface {
	ValidasiUsia(usia int) bool
}

type FormPendaftaran struct {
	NamaLengkap string
	Email       string
	Usia        int
}

func (f FormPendaftaran) ValidasiUsia(usia int) bool {
	if usia < 15 || usia > 75 {
		return false
	}
	return true
}

func ValidasiUsiaForm(formInt FormPendaftaranInterface, usia int) bool {
	return formInt.ValidasiUsia(usia)
}

func inputForm() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nama: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Masukkan usia: ")
	usiaStr, _ := reader.ReadString('\n')
	usiaStr = strings.TrimSpace(usiaStr)

	usia, err := strconv.Atoi(usiaStr)
	if err != nil {
		fmt.Println("Usia harus berupa angka.")
		return
	}

	user := FormPendaftaran{
		NamaLengkap: nama,
		Email:       email,
		Usia:        usia,
	}

	if ValidasiUsiaForm(user, usia) {
		fmt.Println("Pendaftaran berhasil!")
	} else {
		fmt.Println("Usia tidak memenuhi syarat (15 - 75 tahun).")
	}
}

func main() {
	inputForm()
}
