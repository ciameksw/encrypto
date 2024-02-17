package handlers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encrypto/utils"
	"fmt"
	"io"
	"io/ioutil"
)

func EncryptFile(file string, key string) error {
	// Check if file name was provided
	if len(file) == 0 {
		return fmt.Errorf("file parameter is required")
	}

	// Check if key file was provided, if not generate a new one
	var err error
	if len(key) == 0 {
		key, err = utils.GenerateKey()
		if err != nil {
			return fmt.Errorf("generate key err: %v", err.Error())
		}
		fmt.Printf("new key generated: %s\n", key)
	}

	// Read the file data into a variable
	fileData, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("read file err: %v", err.Error())
	}

	// Reading key
	keyFile, err := ioutil.ReadFile(key)
	if err != nil {
		return fmt.Errorf("read key file err: %v", err.Error())
	}

	// Creating block of algorithm
	block, err := aes.NewCipher(keyFile)
	if err != nil {
		return fmt.Errorf("cipher err: %v", err.Error())
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("cipher GCM err: %v", err.Error())
	}

	// Generating random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return fmt.Errorf("nonce err: %v", err.Error())
	}

	// Encrypt file
	encryptedData := gcm.Seal(nonce, nonce, fileData, nil)

	// Generating encrypted file name
	encryptedFileName := utils.GenerateEncryptedFileName(file)

	// Writing encrypted file
	err = ioutil.WriteFile(encryptedFileName, encryptedData, 0777)
	if err != nil {
		return fmt.Errorf("write file err: %v", err.Error())
	}

	// Print encrypted file name
	fmt.Printf("encrypted file created: %s\n", encryptedFileName)
	return nil
}

func DecryptFile(file string, key string) error {
	// Check if file name was provided
	if len(file) == 0 {
		return fmt.Errorf("file parameter is required")
	}

	// Check if key file was provided
	if len(key) == 0 {
		return fmt.Errorf("key parameter is required")
	}

	// Read the encrypted file data into a variable
	encryptedData, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("read file err: %v", err.Error())
	}

	// Reading key
	keyFile, err := ioutil.ReadFile(key)
	if err != nil {
		return fmt.Errorf("read key file err: %v", err.Error())
	}

	// Creating block of algorithm
	block, err := aes.NewCipher(keyFile)
	if err != nil {
		return fmt.Errorf("cipher err: %v", err.Error())
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("cipher GCM err: %v", err.Error())
	}

	// Deattached nonce and decrypt
	nonce := encryptedData[:gcm.NonceSize()]
	encryptedData = encryptedData[gcm.NonceSize():]
	decryptedData, err := gcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return fmt.Errorf("decrypt file err: %v", err.Error())
	}

	// Generating decrypted file name
	decryptedFileName := utils.GenerateDecryptedFileName(file)

	// Writing decryption content
	err = ioutil.WriteFile(decryptedFileName, decryptedData, 0777)
	if err != nil {
		return fmt.Errorf("write file err: %v", err.Error())
	}

	fmt.Printf("decrypted file created: %s\n", decryptedFileName)
	return nil
}
