package helper

func Int64InSlice(val int64, arr []int64) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}

func IsStringInSlice(val string, arr []string) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}
