package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	arr := []string{"0.023", "0.00021", "0.00024", "0.0230", "1.23", "2.0", "2.000", "1", "0.165", "0.164", "0.98"}
	pf := "%-10v %-10v %-10v"
	fmt.Println(fmt.Sprintf(pf, "Input", "", "HUST"))
	for _, a := range arr {
		af, _ := strconv.ParseFloat(strings.TrimSpace(a), 64)
		i1, i2 := Round2SignificantDigits(af)
		istr := fmt.Sprintf("%vx10^%v", i1, i2)
		if i2 == 0 {
			istr = fmt.Sprintf("%v", i1)
		}

		ii1, ii2 := Round2SignificantDigits2(af)
		iistr := fmt.Sprintf("%vx10^%v", ii1, ii2)
		if ii2 == 0 {
			iistr = fmt.Sprintf("%v", ii1)
		}

		fmt.Println(fmt.Sprintf(pf, a, istr, iistr))
	}
}

func Round2SignificantDigits(a float64) (float64, float64) {
	e := math.Floor(math.Log10(a))
	b := math.Pow(10, e)
	x := a / b
	xi, xf := math.Modf(x)
	xf += 2 * 1e-15
	// if b*xf < a/10 {
	// 	return xi, e
	// }
	xfi, _ := math.Modf(xf * 10)
	if x-xi-xfi/10 < xi+(xfi+1)/10-x-2*1e-15 {
		return xi + xfi/10, e
	}
	return xi + (xfi+1)/10, e
}

// en: Rounding errors in physics experiment practice Hanoi University of Science & Technology (HUST)
// vi: Xử lý sai số trong thực hành thí nghiệm vật lý đại học Bách khoa Hà Nội
func Round2SignificantDigits2(a float64) (float64, float64) {
	e := math.Floor(math.Log10(a))
	b := math.Pow(10, e)
	x := a / b
	xi, xf := math.Modf(x)
	xf += 2 * 1e-15
	if b*xf < a/10 {
		return xi, e
	}
	xfi, xff := math.Modf(xf * 10)
	if b*xff/10 < a/10 {
		if x-xi-xfi/10 < xi+(xfi+1)/10-x-2*1e-15 {
			return xi + xfi/10, e
		}
		return xi + (xfi+1)/10, e
	}
	return xi + (xfi+1)/10, e
}
