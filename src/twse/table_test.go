package twse

import (
	"testing"
	"unicode/utf8"
)

func TestConv(t *testing.T) {
	out, err := conv(urlTable)
	if err != nil {
		t.Fatalf("conv exits with error: %v", err)
	}

	if !utf8.Valid(out) {
		t.Error("conv failed")
	}
}
