package utils

import (
	"log"
	"os"
)

func OpenWriteFile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func AppendToFile(path string, entry string) {
	file := OpenWriteFile(path)

	_, err := file.WriteString(entry)
	if err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
