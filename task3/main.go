package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FormPendaftaranInterface interface {
	ValidasiUsia(usia int) bool
	ValidasiGender(gender string) bool
}

type FormLogin struct {
	Username string
	Password string
}

func (f FormLogin) cetakInfoLogin() {
	fmt.Println("Username:", f.Username)
	fmt.Println("Password:", f.Password)
}

func (f *FormLogin) changePassword(password string) {
	f.Password = password
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Masukkan password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	admin1 := FormLogin{
		Username: "admin",
		Password: "123",
	}

	if username != admin1.Username {
		fmt.Println("Username salah. Silahkan coba lagi.")
		return
	}

	if password != admin1.Password {
		fmt.Println("Password salah. Silahkan coba lagi.")
		return
	}

	admin1.cetakInfoLogin()

	fmt.Print("Apakah anda ingin merubah password? (y/n)")
	isChangePassword, _ := reader.ReadString('\n')
	isChangePassword = strings.TrimSpace(strings.ToLower(isChangePassword))

	if isChangePassword == "y" {
		fmt.Print("Masukkan password baru: ")
		newPassword, _ := reader.ReadString('\n')
		newPassword = strings.TrimSpace(newPassword)

		admin1.changePassword(newPassword)
		admin1.cetakInfoLogin()
	}

}
