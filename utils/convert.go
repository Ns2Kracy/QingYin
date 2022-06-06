package utils

import "strconv"

// 将string转换为uint
func StringToUint(str string) (uint, error) {
	i, err := strconv.ParseUint(str, 10, 64)
	return uint(i), err
}
