package strutils

import (
	"testing"
)

func TestToBcryptHash(t *testing.T) {
	str := "every cloud has a silver lining"

	hash, err := ToBcryptHash(str)

	if err != nil {
		t.Error("StrToBcryptHash error: " + err.Error())
	}

	isEqual := BcryptHashCompare(str, hash)
	if isEqual == false {
		t.Error("StrToBcryptHash does not match the comparer")
	}
}
