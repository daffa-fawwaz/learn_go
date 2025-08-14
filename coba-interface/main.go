package main

import (
	"fmt"
	"strings"
)

type FormInterface interface {
	validasiGender(gender string) bool
}

type FormPendaftaran struct {
	NamaLengkap string
	Email       string
	Gender      string
	Usia        int
}

func (f FormPendaftaran) validasiGender(gender string) bool {
	if gender != "laki-laki" {
		return false
	}
	return true
}

func ValidasiGenderForm(formInt FormInterface, gender string) bool {
	return formInt.validasiGender(gender)
}

func main() {
	pendaftar1 := FormPendaftaran{
		NamaLengkap: "Budi Santoso",
		Email:       "email@email.com",
		Gender:      "Laki-laki",
		Usia:        30,
	}

	apakahGenderValid := ValidasiGenderForm(pendaftar1, strings.ToLower(pendaftar1.Gender))
	fmt.Println("Apakah gender valid?", apakahGenderValid)
}
