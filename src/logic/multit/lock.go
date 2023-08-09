package multit

import (
	"fmt"
	"sync"
)

type MyLock struct {
	sync.Mutex
}

type MyRWLock struct {
	sync.RWMutex
}

func Pool() {
	p := sync.Pool{New: func() any {

		return ""
	}}
	p.Put("twj")
	get := p.Get()
	fmt.Println(get)
}
