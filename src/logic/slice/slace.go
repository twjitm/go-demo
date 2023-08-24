package slice

import (
	"fmt"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func GetSlice() {

	slice := []int64{1, 2, 3, 4, 5, 7}
	fmt.Printf("slice %v\n", slice)
	fmt.Printf("slice len%v cap %v\n", len(slice), cap(slice))
	fmt.Printf("slice %p\n", slice)
	changeValue(slice)
	appendValue(slice)
	fmt.Printf("after slice %v\n", slice)
	fmt.Printf("after slice len%v cap %v\n", len(slice), cap(slice))
	fmt.Printf("after slice %p\n", slice)
}

func changeValue(slice []int64) {

	for i, i2 := range slice {
		slice[i] = i2 * 10
	}
}
func appendValue(slice []int64) {

	for i := 0; i < 10; i++ {
		slice = append(slice, int64(i))
	}
	fmt.Printf("target slice %v\n", slice)
	fmt.Printf("target slice len%v cap %v\n", len(slice), cap(slice))
	fmt.Printf("target slice %p\n", slice)
}
