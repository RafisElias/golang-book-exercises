package math

// Calcula a média de uma série de números
func Average(xs []float64) float64 {
	var total float64
	for _, r := range xs {
		total += r
	}
	return total / float64(len(xs))
}
