package hexagon

import (
	"fmt"
	"testing"
)

func TestHex(t *testing.T) {

	for i := 0; i < 10; i++ {
		fmt.Println(i & 1)
	}
}
