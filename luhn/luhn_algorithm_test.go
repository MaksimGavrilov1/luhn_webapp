package luhn

import (
	"testing"
)

func TestValidateEmpty(t *testing.T) {
	if _, err := Validate(""); err == nil {
		t.Fatalf("Validate should not handle empty string")
	}
}
