package hexagon

//odd-q”垂直布局
//将奇数列向下推

type OffsetCoordinate struct {
	Row int //行
	Col int //列
}

func CreMap(maxX, maxY int) []OffsetCoordinate {

	var off []OffsetCoordinate
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			off = append(off, OffsetCoordinate{
				Row: x,
				Col: y,
			})
		}
	}
	return off
}
