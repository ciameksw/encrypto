package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ciameksw/encrypto/handlers"
)

func main() {
	encryptCmd := flag.NewFlagSet("en", flag.ExitOnError)
	encryptFile := encryptCmd.String("f", "", "File to encrypt (required)")
	encryptKey := encryptCmd.String("k", "", "Key used for encryption (optional)")

	decryptCmd := flag.NewFlagSet("de", flag.ExitOnError)
	decryptFile := decryptCmd.String("f", "", "File to decrypt (required)")
	decryptKey := decryptCmd.String("k", "", "Key used for decryption (required)")

	if len(os.Args) < 2 {
		fmt.Println("expected 'en' or 'de' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "en":
		encryptCmd.Parse(os.Args[2:])
		err := handlers.EncryptFile(*encryptFile, *encryptKey)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	case "de":
		decryptCmd.Parse(os.Args[2:])
		err := handlers.DecryptFile(*decryptFile, *decryptKey)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	default:
		fmt.Println("expected 'en' or 'de' subcommands")
		os.Exit(1)
	}
}
