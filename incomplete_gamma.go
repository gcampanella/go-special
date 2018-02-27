package special

import "math"

// The original C code is from http://www.netlib.org/cephes/cprob.tgz

func LowerIncompleteGamma(a, x float64) float64 {
	if x <= 0 || a <= 0 {
		return 0
	}
	if x > 1 && x > a {
		return 1 - UpperIncompleteGamma(a, x)
	}
	ax := a*math.Log(x) - x - lgamma(a)
	if ax < -maxLog {
		return 0
	}
	ax = math.Exp(ax)
	r := a
	c := 1.
	ans := 1.
	for c/ans > machineEpsilon {
		r += 1
		c *= x / r
		ans += c
	}
	return ans * ax / a
}

func UpperIncompleteGamma(a, x float64) float64 {
	const big = 4.503599627370496e15
	const bigInv = 2.22044604925031308085e-16
	if x <= 0 || a <= 0 {
		return 0
	}
	if x < 1 || x < a {
		return 1 - LowerIncompleteGamma(a, x)
	}
	ax := a*math.Log(x) - x - lgamma(a)
	if ax < -maxLog {
		return 0
	}
	ax = math.Exp(ax)
	y := 1 - a
	z := x + y + 1
	c := 0.
	pkm2 := 1.
	qkm2 := x
	pkm1 := x + 1
	qkm1 := z * x
	ans := pkm1 / qkm1
	t := 1.
	for t > machineEpsilon {
		c += 1
		y += 1
		z += 2
		yc := y * c
		pk := pkm1*z - pkm2*yc
		qk := qkm1*z - qkm2*yc
		if qk != 0 {
			r := pk / qk
			t = math.Abs((ans - r) / r)
			ans = r
		} else {
			t = 1
		}
		pkm2 = pkm1
		pkm1 = pk
		qkm2 = qkm1
		qkm1 = qk
		if math.Abs(pk) > big {
			pkm2 *= bigInv
			pkm1 *= bigInv
			qkm2 *= bigInv
			qkm1 *= bigInv
		}
	}
	return ans * ax
}

func lgamma(x float64) (lgamma float64) {
	lgamma, _ = math.Lgamma(x)
	return
}
