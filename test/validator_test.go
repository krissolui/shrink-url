package test

import (
	"shrink-url/internal/util"
	"testing"
)

func TestValidateUrl(t *testing.T) {
	validUrl := "https://example.co"
	if !util.ValidateUrl(validUrl) {
		t.Errorf("ValidateUrl(%s) returned false, should be true", validUrl)
	}

	invalidUrl := "http:/example.co"
	if util.ValidateUrl(invalidUrl) {
		t.Errorf("ValidateUrl(%s) returned true, should be false", invalidUrl)
	}

	invalidUrl = ""
	if util.ValidateUrl(invalidUrl) {
		t.Errorf("ValidateUrl(%s) returned true, should be false", invalidUrl)
	}
}
