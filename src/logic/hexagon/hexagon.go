package hexagon

import (
	"math"
	"unsafe"
)

//odd-q”垂直布局
//将奇数列向下推

// 六边形坐标
type OffsetCoordinate struct {
	Q int32 //列
	R int32 //行

}

func (c *OffsetCoordinate) HashCode() int64 {
	return *(*int64)(unsafe.Pointer(c))
}

func (c *OffsetCoordinate) IsNull() bool {
	return c.Q == 0 && c.R == 0
}

func (c *OffsetCoordinate) Clear() {
	c.Q, c.R = 0, 0
}

// 二维坐标系,平面坐标系中计算距离
type DoubleCoordinate struct {
	X int32
	Y int32
}

func (c *DoubleCoordinate) IsNullOrZero() bool {
	return c.X == 0 && c.Y == 0
}

// FCoordinate 平面直角坐标系，浮点数坐标，用于向量计算
type FCoordinate struct {
	X float64
	Y float64
}

// Skew 求法向量
func (f FCoordinate) Skew() FCoordinate {
	return FCoordinate{X: -f.Y, Y: f.X}
}

// Normalize 向量归一化
func (f FCoordinate) Normalize() FCoordinate {
	dis := math.Sqrt(f.X*f.X + f.Y*f.Y)
	return FCoordinate{
		X: f.X / dis,
		Y: f.Y / dis,
	}
}

// CubeCoordinate cube坐标表示
type CubeCoordinate struct {
	X int32
	Y int32
	Z int32
}

// FCubeCoordinate cube浮点坐标，一般用于线性插值计算
type FCubeCoordinate struct {
	X float64
	Y float64
	Z float64
}

// cubeRound 线性插值法计算线段经过哪些tile时，需要把float坐标转成最近的hex tile
func cubeRound(cube FCubeCoordinate) CubeCoordinate {
	var rx = math.Round(cube.X)
	var ry = math.Round(cube.Y)
	var rz = math.Round(cube.Z)

	var xDiff = math.Abs(rx - cube.X)
	var yDiff = math.Abs(ry - cube.Y)
	var zDiff = math.Abs(rz - cube.Z)

	if xDiff > yDiff && xDiff > zDiff {
		rx = -ry - rz
	} else if yDiff > zDiff {
		ry = -rx - rz
	} else {
		rz = -rx - ry
	}

	return CubeCoordinate{
		X: int32(rx),
		Y: int32(ry),
		Z: int32(rz),
	}
}

// cubeRound 线性插值法计算线段经过哪些tile时，需要把float坐标转成最近的hex tile
func cubeRoundRayCast(cube FCubeCoordinate) CubeCoordinate {
	var rx = math.Round(cube.X)
	var ry = math.Round(cube.Y)
	var rz = math.Round(cube.Z)

	var xDiff = math.Abs(rx - cube.X)
	var yDiff = math.Abs(ry - cube.Y)
	var zDiff = math.Abs(rz - cube.Z)

	if xDiff > yDiff && xDiff > zDiff {
		rx = -ry - rz
	} else if yDiff >= zDiff {
		ry = -rx - rz
	} else {
		rz = -rx - ry
	}

	return CubeCoordinate{
		X: int32(rx),
		Y: int32(ry),
		Z: int32(rz),
	}
}

// DoubleToCube double坐标转cube坐标
func DoubleToCube(c DoubleCoordinate) CubeCoordinate {
	x := c.X
	z := (c.Y - c.X) / 2
	y := -x - z
	return CubeCoordinate{X: x, Y: y, Z: z}
}

// CubeToDouble cube坐标转double坐标
func CubeToDouble(c CubeCoordinate) DoubleCoordinate {
	x := c.X
	y := 2*c.Z + c.X
	return DoubleCoordinate{X: x, Y: y}
}

// DoubleDistance double坐标距离
func DoubleDistance(c1, c2 DoubleCoordinate) int32 {
	dx := Abs(c2.X - c1.X)
	dy := Abs(c2.Y - c1.Y)
	return dx + Max(0, (dy-dx)/2)
}

// CubeDistance cube坐标距离
func CubeDistance(c1, c2 CubeCoordinate) int32 {
	return (Abs(c1.X-c2.X) + Abs(c1.Y-c2.Y) + Abs(c1.Z-c2.Z)) / 2
}

// CubeToOffset cube坐标转成offset坐标
func CubeToOffset(c CubeCoordinate) OffsetCoordinate {
	q := c.X
	r := c.Z + (c.X-(c.X&1))/2
	return OffsetCoordinate{Q: q, R: r}
}

// OffsetToCube offset坐标转cube坐标
func OffsetToCube(c OffsetCoordinate) CubeCoordinate {
	x := c.Q
	z := c.R - (c.Q-(c.Q&1))/2
	y := -x - z
	return CubeCoordinate{X: x, Y: y, Z: z}
}

// OffsetToDouble offset坐标转成double坐标
// 把Q给x，若Q为奇数，则y=2*R+1，否则y=2R
func OffsetToDouble(c OffsetCoordinate) DoubleCoordinate {
	x := c.Q
	parity := c.Q & 1
	y := 2*c.R + parity
	return DoubleCoordinate{X: x, Y: y}
}

// DoubleToOffset double坐标转成offset坐标
func DoubleToOffset(c DoubleCoordinate) OffsetCoordinate {
	return OffsetCoordinate{Q: c.X, R: c.Y / 2}
}

// FCoordinateToOffset 平面直角坐标系转offset
func FCoordinateToOffset(c FCoordinate) OffsetCoordinate {
	q := math.Sqrt(3)/3*c.X - 1./3*c.Y
	r := 2. / 3 * c.Y
	fCube := FCubeCoordinate{X: q, Y: -q - r, Z: r}
	return CubeToOffset(cubeRound(fCube))
}

// OffsetToFCoordinate Offset转成平面直角坐标系
func OffsetToFCoordinate(c OffsetCoordinate) FCoordinate {
	dCoord := OffsetToDouble(c)
	y := math.Sqrt(3) / 2 * float64(dCoord.Y)
	x := float64(3) / 2 * float64(dCoord.X)
	return FCoordinate{X: x, Y: y}
}

// FCoordDistance 平面直角坐标系下两点间直线距离
func FCoordDistance(c1, c2 FCoordinate) float64 {
	dx := c2.X - c1.X
	dy := c2.Y - c1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// OffsetDistance 两个offset坐标的六边形距离
func OffsetDistance(c1, c2 OffsetCoordinate) int32 {
	f1 := OffsetToCube(c1)
	f2 := OffsetToCube(c2)
	return CubeDistance(f1, f2)
}
