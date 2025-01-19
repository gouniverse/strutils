package strutils

import (
	"bytes"
	"encoding/base64"
)

// Base64Decode decodes base64 text to binary data.
func Base64Decode(text []byte) (data []byte, err error) {
	n := base64.StdEncoding.DecodedLen(len(text))
	data = make([]byte, n)
	decoder := base64.NewDecoder(base64.StdEncoding, bytes.NewBuffer(text))
	if n, err = decoder.Read(data); err != nil {
		n = 0
	}
	data = data[:n]
	return
}
