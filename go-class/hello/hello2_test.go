package hello

import (
	"fmt"
	"testing"
)

func TestSayHello(t *testing.T) {
	want := "Hello, test!"
	got := Say("test")
	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}
	fmt.Println("wanted ", want, " got ", got)
}
