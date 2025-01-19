package strutils

import "unsafe"

// StrToBytes converts string to bytes
func ToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
