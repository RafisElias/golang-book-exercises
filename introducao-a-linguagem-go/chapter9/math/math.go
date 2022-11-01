package math

// Average - Calcula a média de uma série de números
func Average(xs []float64) float64 {
	var total float64
	if len(xs) == 0 {
		return 0
	}
	for _, r := range xs {
		total += r
	}
	return total / float64(len(xs))
}

// Min - Retorna o menor valor de um slice
func Min(xs []float64) float64 {
	var min float64
	for i, r := range xs {
		if i == 0 || r <= min {
			min = r
		}
	}
	return min
}

// Max - Retorna o maior valor de um slice
func Max(xs []float64) float64 {
	var max float64
	for i, r := range xs {
		if i == 0 || r >= max {
			max = r
		}
	}
	return max
}
