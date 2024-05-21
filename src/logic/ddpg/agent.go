package ddpg

import "math"

func optimal_f(p float64, pi float64, lambd float64, psi float64, cost string) float64 {

	if cost == "trade_0" {
		return p/(2*lambd) - pi
	}
	if cost == "trade_l2" {
		return p/(2*(lambd+psi)) + psi*pi/(lambd+psi) - pi
	}
	if cost == "trade_l1" {
		if p <= -psi+2*lambd*pi {
			return (p+psi)/(2*lambd) - pi
		}
		if (-psi+2*lambd*pi < p) && (p < psi+2*lambd*pi) {
			return 0
		}
		if p >= psi+2*lambd*pi {
			return (p-psi)/(2*lambd) - pi
		}

	}
	return 0
}

func optimal_max_pos(p float64, pi float64, thresh float64, maxPos float64) float64 {

	if math.Abs(p) < thresh {
		return 0
	}
	if p > thresh {
		return maxPos - pi
	}
	if p <= -thresh {
		return -maxPos - pi
	}

	return 0
}
