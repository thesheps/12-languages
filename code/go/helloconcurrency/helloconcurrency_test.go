package helloconcurrency

import (
	"fmt"
	"testing"
)

// TestHello our HelloWorld functionality
func TestHello(t *testing.T) {
	message := Hello()
	
	if message != "Hello World!" {
		t.Error(fmt.Printf("Message didn't equal HelloWorld, equalled %s", message))
	}
}