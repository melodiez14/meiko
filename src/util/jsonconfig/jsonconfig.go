package jsonconfig

import (
	"encoding/json"
	"fmt"
	"os"
)

func Load(cfg interface{}, path string, env string) bool {
	p := fmt.Sprintf("%s/%s.json", path, env)
	file, err := os.Open(p)
	if err != nil {
		return false
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(cfg)
	if err != nil {
		return false
	}
	return true
}
