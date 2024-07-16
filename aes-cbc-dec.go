package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	openssl "github.com/Luzifer/go-openssl/v4"

	"golang.org/x/term"
)

func main() {
	fmt.Print("Password: ")
	bytepw, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	passphrase := string(bytepw)

	data_b64, err := os.ReadFile("./config_enc_b64.cnf")
	if err != nil {
		log.Fatal(err)
	}
	println(string(data_b64))

	// data, err := b64.StdEncoding.DecodeString(string(data_b64))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	o := openssl.New()

	decrypted, err := o.DecryptBytes(passphrase, data_b64, openssl.PBKDF2SHA256)
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
	}
	fmt.Printf("Decrypted text: %s\n", string(decrypted))

}
