package strutils

import (
	"bytes"
	"encoding/base32"
)

// Base32ExtendedDecode encodes binary data to base32 extended (RFC 4648) encoded text.
func Base32ExtendedDecode(data []byte) (text []byte, err error) {
	n := base32.StdEncoding.DecodedLen(len(text))
	data = make([]byte, n)
	decoder := base32.NewDecoder(base32.StdEncoding, bytes.NewBuffer(text))

	if n, err = decoder.Read(data); err != nil {
		n = 0
	}

	text = data[:n]

	return text, nil
}
