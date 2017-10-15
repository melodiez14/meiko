package alias

import (
	"log"
	"os"
)

// DirectoryConfig is loaded from json configuration
type DirectoryConfig struct {
	Static  string `json:"static"`
	Profile string `json:"profile"`
	Email   string `json:"email"`
}

// Dir is the list of directory
var Dir DirectoryConfig

// InitDirectory is used for initialize all of file directory
func InitDirectory(cfg DirectoryConfig) {
	log.Println("Initializing directory")

	if _, err := os.Stat(cfg.Static); os.IsNotExist(err) {
		log.Fatalf("Static: %s is not exist", cfg.Static)
		return
	}

	if _, err := os.Stat(cfg.Profile); os.IsNotExist(err) {
		log.Fatalf("Profile: %s is not exist", cfg.Profile)
		return
	}

	if _, err := os.Stat(cfg.Email); os.IsNotExist(err) {
		log.Fatalf("Email: %s is not exist", cfg.Email)
		return
	}

	Dir = cfg

	log.Println("Directory successfully loaded")
}
