package guess

//y=mx+b
// LinearRegression calculates the linear regression line coefficients (slope and y-intercept)
// for a given set of y-values. The x-values are assumed to be the indices of the y-values.
//
// The function takes a slice of float64 values representing the y-values.
//
// It returns two float64 values:
// - The first value is the slope (m) of the regression line.
// - The second value is the y-intercept (b) of the regression line.
func LinearRegression(values []float64) (float64, float64) {
	n := float64(len(values))
	var sumX, sumY, sumXY, sumX2 float64

	for i, y := range values {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
	}
	// m = N(∑x2) - (∑x)² / N(∑xy) - (∑x)(∑y)
	m := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	// b = N(∑y) - m(∑x) / N
	b := (sumY - m*sumX) / n

	return m, b
}
