package alias

import (
	"log"
	"os"
)

// DirectoryConfig is loaded from json configuration
type DirectoryConfig struct {
	Public  string `json:"public"`
	Profile string `json:"profile"`
}

// Dir is the list of directory
var Dir DirectoryConfig

func InitDirectory(cfg DirectoryConfig) {
	log.Println("Initializing directory")

	if _, err := os.Stat(cfg.Public); os.IsNotExist(err) {
		log.Fatalf("Public: %s is not exist", cfg.Public)
		return
	}

	if _, err := os.Stat(cfg.Profile); os.IsNotExist(err) {
		log.Fatalf("Profile: %s is not exist", cfg.Profile)
		return
	}

	Dir = cfg

	log.Println("Directory successfully loaded")
}
