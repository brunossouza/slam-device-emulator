package utils

import (
	"log"
	"os"
)

// CheckError verifica erros
func CheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// FileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
