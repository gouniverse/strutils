package strutils

import (
	"encoding/base32"
)

// Base32ExtendedDecode encodes binary data to base32 extended (RFC 4648) encoded text.
func Base32ExtendedDecode(data []byte) (text []byte, err error) {
	return base32.StdEncoding.DecodeString(string(data))
}
