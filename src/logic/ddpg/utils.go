package ddpg

import (
	"math"
	"math/rand"
)

const DefaultRandomState = 1000

func Build_ou_process(T int64, theta float64, sigma float64, randomState int32) []float64 {

	var t = 0
	var x = 0.0
	if randomState == 0 {
		randomState = DefaultRandomState
	}
	rand.New(rand.NewSource(int64(randomState)))
	normals := make([]float64, T)
	for i := int64(0); i < T; i++ {
		u1 := rand.Float64()
		u2 := rand.Float64()
		n := math.Sqrt(-2*math.Log(u1)) * math.Cos(2*math.Pi*u2) // N(0, 1)
		normals[i] = n
	}
	X := make([]float64, T)
	for i, normal := range normals {
		x += -x*theta + sigma*normal
		var d = sigma * math.Sqrt(1.0/2.0/theta) //缩放处理
		x = x / d
		X[i] = x
	}
	x += -x*theta + sigma*normals[t]
	return X
}

func GetReturns(signal []float64, randomState int32) []float64 {

	if randomState == 0 {
		randomState = DefaultRandomState
	}
	var rand2 = rand.New(rand.NewSource(int64(randomState)))
	X := make([]float64, len(signal))
	for i, f := range signal {
		v := rand2.Float64()
		X[i] = f + v
	}
	return X
}
