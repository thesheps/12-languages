package hellohttp

import (
	"testing"
)

// TestHello tests our HelloWorld functionality
func TestHello(t *testing.T) {
	message := Hello()
	
	if message != "Hello World!" {
		t.Error("Message didn't equal HelloWorld :(")
	}
}