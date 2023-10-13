package strutils

import (
	"testing"
)

func TestBase32ExtendedEncode(t *testing.T) {
	text := "Hello world"

	textEnc, err := Base32ExtendedEncode([]byte(text))

	if err != nil {
		t.Error("Error", err.Error())
		t.Fail()
	}

	if string(textEnc) != "JBSWY3DPEB3W64TMMQ======" {
		t.Error("Enc text must be equal to JBSWY3DPEB3W64TMMQ====== but found:", string(textEnc))
	}

	textDec, err := Base32ExtendedDecode([]byte(textEnc))

	if err != nil {
		t.Error("Error:", err.Error())
		t.Fail()
	}

	if string(textDec) != "123456" {
		t.Error("Dec text must be equal to 123456 but found:", string(textDec))
	}
}
