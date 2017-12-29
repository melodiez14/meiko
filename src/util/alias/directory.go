package alias

import (
	"log"
	"os"
)

// DirectoryConfig is loaded from json configuration
type DirectoryConfig map[string]string

// Dir is the list of directory
var Dir DirectoryConfig

// InitDirectory is used for initialize all of file directory
func InitDirectory(cfg DirectoryConfig) {
	log.Println("Initializing directory")

	for i, val := range cfg {
		if _, err := os.Stat(val); os.IsNotExist(err) {
			log.Fatalf("%s: %s is not exist", i, val)
			return
		}
	}

	Dir = cfg

	log.Println("Directory successfully loaded")
}
