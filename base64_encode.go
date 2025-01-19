package strutils

import (
	"bytes"
	"encoding/base64"
	"errors"
)

// Base64Encode encodes binary data to base64 encoded text.
func Base64Encode(data []byte) (text []byte, err error) {
	n := base64.StdEncoding.EncodedLen(len(data))
	buf := bytes.NewBuffer(make([]byte, 0, n))
	encoder := base64.NewEncoder(base64.StdEncoding, buf)
	encoder.Write(data)
	encoder.Close()
	if buf.Len() != n {
		return text, errors.New("zero buffer length")
	}
	return buf.Bytes(), nil
}
