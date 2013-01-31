package foo

import (
	"testing"
)

var squareTests = []struct {
	num, square Number
}{
	{1, 1},
	{2, 4},
	{256, 65536},
	{-10, 100},
}

func TestSquare(t *testing.T) {
	for _, test := range squareTests {
		actual := test.num.Square()
		if actual != test.square {
			t.Errorf("Square() of %v: got %v, want %v",
				test.num, actual, test.square)
		}
	}
}
