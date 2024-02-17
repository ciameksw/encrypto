package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func GenerateKey() (string, error) {
	// Generating key
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return "", err
	}

	// Generate file name
	fileName := fmt.Sprintf("key-%v.txt", time.Now().Unix())

	// Writing key file
	err := ioutil.WriteFile(fileName, key, 0777)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func GenerateEncryptedFileName(file string) string {
	// Get only file name from path
	file = filepath.Base(file)

	// Add timestamp and .encrypted.bin extension to the file name
	return fmt.Sprintf("%s-%v.encrypted.bin", file, time.Now().Unix())
}

func GenerateDecryptedFileName(file string) string {
	// Get only file name from path
	file = filepath.Base(file)

	// Remove possible encrypted file extension
	file = strings.TrimSuffix(file, ".encrypted.bin")
	file = strings.TrimSuffix(file, ".bin")

	// Check if file has an extension
	r := regexp.MustCompile(`^[^.]+\.[\w]+`)
	matched := r.FindString(file)

	// If file has an extension, generate a new name with the same extension
	if matched != "" {
		ext := filepath.Ext(matched)
		name := strings.TrimSuffix(matched, ext)
		return fmt.Sprintf("%s-%v.decrypted%s", name, time.Now().Unix(), ext)
	}

	// If file has no extension, generate a new name
	// with only alphanumeric characters from the beginning of the file name
	r = regexp.MustCompile(`^[\w]+`)
	matched = r.FindString(file)
	return fmt.Sprintf("%s-%v.decrypted", matched, time.Now().Unix())
}
