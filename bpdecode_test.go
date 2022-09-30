package bpencode_test

import (
	"testing"

	"github.com/waldirborbajr/bpencode"
)

// https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318
func TestBPEncodeEmpty(t *testing.T) {
	str, err := bpencode.BPEncode("")

	if err != nil {
		t.Logf("error empty password found %v", err)
	} else {
		t.Logf("success: %v", str)
	}
}

func TestBPEncodeSimple(t *testing.T) {
	str, err := bpencode.BPEncode("bataca")

	if err != nil {
		t.Errorf("error empty password %v", err)
	} else {
		t.Logf("success: %v", str)
	}
}

func TestBPEncodeComplex(t *testing.T) {
	str, err := bpencode.BPEncode(":/?#[]@senha()2bataca")
	if err != nil {
		t.Errorf("error empty password %v", err)
	} else {
		t.Logf("success: %v", str)
	}
}
