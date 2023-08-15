package lru

import (
	"fmt"
	"strconv"
	"testing"
)

func TestLru(t *testing.T) {
	cache := NewLRUCache(10)
	for i := 0; i < 10; i++ {
		cache.Add("key_"+strconv.Itoa(i), i)
	}
	value := cache.Get("key_5")
	cache.Add("key_11", 11)
	fmt.Println(value)
}
