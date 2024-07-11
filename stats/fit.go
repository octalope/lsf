package stats

import (
	"math"
)

type Sums struct {
    x, y, xx, xy, yy float64
}

func calculateSums(records [][]float64, xIndex int, yIndex int) Sums {
    var x, y, xx, xy, yy float64
    for _, record := range records {
        x += record[xIndex]
        y += record[yIndex]
        xx += record[xIndex] * record[xIndex]
        xy += record[xIndex] * record[yIndex]
        yy += record[yIndex] * record[yIndex]
    }

    return Sums{x, y, xx, xy, yy}
}

func LeastSquaresFit(records [][]float64, xIndex int, yIndex int) (float64, float64, float64, float64, float64) {
    sums := calculateSums(records, xIndex, yIndex)

    n := float64(len(records))
    m := (n * sums.xy - sums.x * sums.y) / (n * sums.xx - sums.x * sums.x)
    b := (sums.y - m * sums.x) / n
    rSquared := (sums.xy * sums.xy) / (sums.xx * sums.yy)
    standardDeviationSquared := (sums.yy - (sums.xy * sums.xy / sums.xx)) / (n - 2)
    standardDeviation := math.Sqrt(math.Abs(standardDeviationSquared))
    meanX := sums.x / n
        
    dm := standardDeviation * math.Sqrt((1.0 / n) + (meanX * meanX / sums.xx))
    db := standardDeviation / math.Sqrt(sums.xx)

    return m, dm, b, db, rSquared
}
