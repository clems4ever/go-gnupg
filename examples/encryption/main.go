package main

import (
	"context"
	"fmt"
	"log"

	gognupg "github.com/clems4ever/go-gnupg"
)

func main() {
	gpg := gognupg.NewGnuPG()

	var yourEmail string
	fmt.Print("Email: ")
	_, err := fmt.Scan(&yourEmail)
	if err != nil {
		log.Panic(err)
	}

	cipher, err := gpg.Encrypt(context.Background(), []byte("secure-data"), []string{yourEmail}, &gognupg.EncryptionOptions{
		Armor: true,
	})
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("gnupg generated the armor cipher:\n", string(cipher))

	plaintext, err := gpg.Decrypt(context.Background(), cipher)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("decrypted plain text: ", string(plaintext))
}
