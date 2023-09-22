package test

import (
	"shrink-url/internal/util"
	"testing"
)

func TestEncryptUrl(t *testing.T) {
	input := "https://www.quantbe.com/welcome/canada/logs/validate"
	maxLenth := 6
	encryptor := util.NewCryptography()
	output := encryptor.EncryptUrl(input, maxLenth)

	if len(output) > maxLenth {
		t.Errorf("output length is %d, max length should be %d", len(output), maxLenth)
	}

	if output != "ysrAXm" {
		t.Errorf("EncryptUrl(%s, %d) returned %s, should be ysrAXm", input, maxLenth, output)
	}
}
