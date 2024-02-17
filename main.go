package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ciameksw/encrypto/handlers"
)

func main() {
	encryptCmd := flag.NewFlagSet("encrypt", flag.ExitOnError)
	encryptFile := encryptCmd.String("f", "", "File to encrypt (required)")
	encryptKey := encryptCmd.String("k", "", "Key used for encryption (optional)")

	decryptCmd := flag.NewFlagSet("decrypt", flag.ExitOnError)
	decryptFile := decryptCmd.String("f", "", "File to decrypt (required)")
	decryptKey := decryptCmd.String("k", "", "Key used for decryption (required)")

	if len(os.Args) < 2 {
		fmt.Println("expected 'encrypt' or 'decrypt' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "encrypt":
		encryptCmd.Parse(os.Args[2:])
		err := handlers.EncryptFile(*encryptFile, *encryptKey)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	case "decrypt":
		decryptCmd.Parse(os.Args[2:])
		err := handlers.DecryptFile(*decryptFile, *decryptKey)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	default:
		fmt.Println("expected 'encrypt' or 'decrypt' subcommands")
		os.Exit(1)
	}
}
