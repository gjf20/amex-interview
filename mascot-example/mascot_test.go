package mascotexample_test

import (
	"testing"

	mascotexample "example.com/go-amex-interview/mascot-example"
)

func TestMascot(t *testing.T) {
	if mascotexample.Mascot() != "Go Gopher" {
		t.Fatal("wrong mascot")
	}
}
