package minitask1

import "math"

func Luas(r float32) float32 {
	return math.Pi * float32(r) * float32(r)
}

func Keliling(r float32) float32 {
	return 2 * math.Pi * float32(r)
}

func LuasDanKeliling(r float32) (luas float32, keliling float32) {
	luas = Luas(r)
	keliling = Keliling(r)
	return luas, keliling
}
