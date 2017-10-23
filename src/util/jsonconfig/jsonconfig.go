package jsonconfig

import (
	"encoding/json"
	"fmt"
	"os"
)

// Load loads json file succes or not 
/*
	@params:
		cfg		= interface{}
		path	= string
		env		= string
	@example:
		cfg		= interface{}
		path	= c:/file
		env		= development
	@return
*/
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
