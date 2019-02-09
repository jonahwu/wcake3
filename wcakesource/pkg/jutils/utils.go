package jutils

import (
	"fmt"
	"github.com/satori/go.uuid"
	"os"
)

func GetHostName() string {
	hn, _ := os.Hostname()
	return hn

}

func GetUuid() string {
	return uuid.Must(uuid.NewV4()).String()
}

func CombineString(str1 string, str2 string) string {
	return fmt.Sprintf("%s%s", str1, str2)
}
