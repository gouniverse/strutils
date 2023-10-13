package strutils

import "strconv"

func IntToBase36(num int) string {
	return strconv.FormatInt(int64(num), 36)
}
