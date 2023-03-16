package strutils

import (
	"bytes"
	"encoding/base64"
)

// Base64Encode encodes binary data to base64 encoded text.
func Base64Encode(data []byte) (text []byte) {
	n := base64.StdEncoding.EncodedLen(len(data))
	buf := bytes.NewBuffer(make([]byte, 0, n))
	encoder := base64.NewEncoder(base64.StdEncoding, buf)
	encoder.Write(data)
	encoder.Close()
	if buf.Len() != n {
		panic("internal error")
	}
	return buf.Bytes()
}
