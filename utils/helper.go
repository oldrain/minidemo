package utils

import (
	"encoding/json"
	"fmt"
	"time"
	"unicode"
)

const (
	DateTimeLayout24 = "2006-01-02 15:04:05"
	DateLayout = "2006-01-02"
	DateLayoutClose = "20060102"
)

func ObjToJson(s interface{}) string {
	js, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(js)
}

func JsonToObj(s string, obj interface{}) (err error) {
	err = json.Unmarshal([]byte(s), obj)
	return
}

func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func FormatString(s string, args ...interface{}) string {
	return fmt.Sprintf(s, args...)
}

func NowDateTime() string {
	return time.Now().Format(DateTimeLayout24)
}

func NowDateClose() string {
	return time.Now().Format(DateLayoutClose)
}

func NowDate() string {
	return time.Now().Format(DateLayout)
}
