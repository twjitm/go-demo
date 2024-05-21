package patchfinding

import (
	"context"
	"go-demo/src/logic/hexagon"
)

//寻路代价

func Cost(ctx context.Context, src, dest hexagon.OffsetCoordinate) float64 {
	return euclideanCost(ctx, src, dest)
}

func euclideanCost(ctx context.Context, src, dest hexagon.OffsetCoordinate) float64 {
	s := hexagon.OffsetToDouble(src)
	d := hexagon.OffsetToDouble(dest)
	if s == d {
		return 0
	}
	//vector := hexagon.DoubleCoordinate{X: d.X - s.X, Y: d.Y - s.Y} // 平移到0,0
	//dir1, dir2 := hexagon.FindNearestDirection(ctx, vector)
	//if dir1 == dir2 {
	//	var scale float64
	//	if dir1.X != 0 {
	//		scale = float64(vector.X / dir1.X)
	//	} else {
	//		scale = float64(vector.Y / dir1.Y)
	//	}
	//	return calcDistanceInDirectionX(dir1, scale)
	//}
	//// 如果dir1和dir2不是同一个方向，那对于src来说一定是一个sibling一个cousin方向
	//k := dir1.Y*dir2.X - dir1.X*dir2.Y
	//n := (dir1.Y*vector.X - dir1.X*vector.Y) / k
	//m := (dir2.X*vector.Y - dir2.Y*vector.X) / k
	//res := calcDistanceInDirectionX(dir1, float64(m))
	//res += calcDistanceInDirectionX(dir2, float64(n))
	//return res
	return 0
}

func manhattanCost(src, dest hexagon.OffsetCoordinate) float64 {
	c1 := hexagon.OffsetToCube(src)
	c2 := hexagon.OffsetToCube(dest)
	return float64(hexagon.CubeDistance(c1, c2))
}
