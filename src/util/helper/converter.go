package helper

import (
	"crypto/md5"
	"fmt"
)

func StringToMD5(text string) string {
	data := []byte(text)
	return fmt.Sprintf("%x", md5.Sum(data))
}
