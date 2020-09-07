package util

import "strconv"

func CvtIntToString(i int) string {
	return strconv.Itoa(i)
}

func CvtStringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func CvtInt64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func CvtStringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}
