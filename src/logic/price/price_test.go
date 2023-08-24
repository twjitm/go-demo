package price

import (
	"fmt"
	"testing"
)

func TestPrice(t *testing.T) {
	var slices []int64
	slices = append(slices, 7)
	slices = append(slices, 1)
	slices = append(slices, 3)
	slices = append(slices, 5)
	slices = append(slices, 3)
	slices = append(slices, 6)
	v := getMaxPrice(slices)
	fmt.Println(v)

}
