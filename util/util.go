package util

import (
	"log"
	"os"
)

// GetUserDir returns the path to the user's home directory.
func GetUserDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return home
}
