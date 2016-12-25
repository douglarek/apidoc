package main

import (
	"testing"
)

// I am a func doc.
func _(_ string, _ bool) {}

func TestParseDocs(t *testing.T) {
	docs, err := parseDocs(".", []string{"string", "bool"})
	if err != nil {
		t.Error(err)
	}
	if actual, expect := string(docs), "I am a func doc.\n"; actual != expect {
		t.Errorf("Expected `%v`, but it was `%v` instead.", expect, actual)
	}
}
