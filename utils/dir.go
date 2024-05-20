package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func GetDirFromHome(path ...string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	fullPath := append([]string{home}, path...)

	dir := filepath.Join(fullPath...)
	return dir
}

func GetConfigDir() string {
	return GetDirFromHome(".config", "prt")
}

func GetConfigFile() string {
  return GetDirFromHome(".config", "prt", "config.yaml")
}

func GetCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func GetDirFromCurr(path ...string) string {
	dir := GetCurrentDir()

	fullPath := append([]string{dir}, path...)

	newDir := filepath.Join(fullPath...)
	return newDir
}

func DirExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}
