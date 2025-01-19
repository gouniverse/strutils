package strutils

import (
	"bytes"
	"encoding/base32"
	"errors"
)

// Base32ExtendedEncode encodes binary data to base32 extended (RFC 4648) encoded text.
func Base32ExtendedEncode(data []byte) (text []byte, err error) {
	n := base32.StdEncoding.EncodedLen(len(data))
	buf := bytes.NewBuffer(make([]byte, 0, n))
	encoder := base32.NewEncoder(base32.StdEncoding, buf)
	encoder.Write(data)
	encoder.Close()

	if buf.Len() != n {
		return text, errors.New("zero buffer length")
	}

	return buf.Bytes(), nil
}
