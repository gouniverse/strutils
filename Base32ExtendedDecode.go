package strutils

import (
	"bytes"
	"encoding/base32"
)

// Base32ExtEncode encodes binary data to base32 extended (RFC 4648) encoded text.
func Base32ExtEncode(data []byte) (text []byte) {
	n := base32.HexEncoding.EncodedLen(len(data))
	buf := bytes.NewBuffer(make([]byte, 0, n))
	encoder := base32.NewEncoder(base32.HexEncoding, buf)
	encoder.Write(data)
	encoder.Close()
	if buf.Len() != n {
		panic("internal error")
	}
	return buf.Bytes()
}
