package patchfinding

import "go-demo/src/logic/hexagon"

type Node struct {
	parent *Node
	G      float64
	H      float64
	F      float64
	index  int
	hexagon.OffsetCoordinate
	direction *hexagon.DoubleCoordinate // parent node到当前node的方向
}
