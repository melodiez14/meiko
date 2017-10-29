package helper

// Int64InSlice Check if the value given are in part of array value in Int format
/*
	@params:
		val	= int
		arr = []int
	@example:
		val = 64
		arr	= [64,44,123,343]
	@return
		true/false
*/
func Int64InSlice(val int64, arr []int64) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}

// Int8InSlice Check if the value given are in part of array value in Int format
/*
	@params:
		val	= int
		arr = []int
	@example:
		val = 64
		arr	= [64,44,123,343]
	@return
		true/false
*/
func Int8InSlice(val int8, arr []int8) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}

// IsStringInSlice Check if the value given are in part of array value in String format
/*
	@params:
		val	= String
		arr = []String
	@example:
		val = me
		arr	= [me,you,he,she,they]
	@return
		true/false
*/
func IsStringInSlice(val string, arr []string) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}
