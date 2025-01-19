package strutils

import (
	"strconv"
)

func IntToBase32(num int) string {
	return strconv.FormatInt(int64(num), 32)
}
