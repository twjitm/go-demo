package price

import (
	"fmt"
	"testing"
)

func TestPrice(t *testing.T) {
	var slices = []int64{1, 2, 3, 4, 5}
	v := getMaxPrice(slices)
	fmt.Println(v)

}
