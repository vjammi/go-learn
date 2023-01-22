package sumup

import (
	"testing"
)

func TestSum(t *testing.T) {
	reference := 40 + 2
	result := Sum(40, 2)

	if result != reference {
		t.Errorf("Sum should return %d instead of %d", reference, result)
	}
}
