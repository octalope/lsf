package stats

import (
	"math"
)

func calculateSums(records [][]float64, xIndex int, yIndex int) (float64, float64, float64, float64, float64) {
    var x, y, xx, xy, yy float64
    for _, record := range records {
        x += record[xIndex]
        y += record[yIndex]
        xx += record[xIndex] * record[xIndex]
        xy += record[xIndex] * record[yIndex]
        yy += record[yIndex] * record[yIndex]
    }

    return x, y, xx, xy, yy
}

func LeastSquaresFit(records [][]float64, xIndex int, yIndex int) (float64, float64, float64, float64, float64) {
    x, y, xx, xy, yy := calculateSums(records, xIndex, yIndex)

    n := float64(len(records))
    m := (n*xy - x*y) / (n*xx - x*x)
    b := (y - m*x) / n
    rSquared := (xy * xy) / (xx * yy)
    standardDeviationSquared := (yy - (xy * xy / xx)) / (n - 2)
    standardDeviation := math.Sqrt(math.Abs(standardDeviationSquared))
    meanX := x / n
        
    dm := standardDeviation * math.Sqrt((1.0 / n) + (meanX * meanX / xx))
    db := standardDeviation / math.Sqrt(xx)

    return m, dm, b, db, rSquared
}
