package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Nina")
	if s != "Welcome my dear Nina" {
		t.Error("got - ", s, "expected - ", "Welcome my dear Nina")
	}
}
func ExampleGreet() {
	fmt.Println(Greet("Nina"))
	// Output:
	// Welcome my dear Nina
}

func BenchmartGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Nina")
	}
}
