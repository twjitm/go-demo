package hexagon

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestHex(t *testing.T) {

	for i := 0; i < 10; i++ {
		fmt.Println(i & 1)
	}
}

func TestCoord(t *testing.T) {

	for i := 0; i < 5; i++ {

		start := OffsetCoordinate{
			R: rand.Int31n(1024),
			Q: rand.Int31n(1024),
		}
		end := OffsetCoordinate{
			R: rand.Int31n(1024),
			Q: rand.Int31n(1024),
		}
		fmt.Println("start r q", start.R, start.Q)
		s := OffsetToDouble(start)
		fmt.Println("start double s r q", s.X, s.Y)
		fmt.Println("---------------------------------")
		fmt.Println("end r q", end.R, end.Q)
		d := OffsetToDouble(end)
		fmt.Println("end double d  r q", d.X, d.Y)
		fmt.Println("---------------------------------")
	}

}
