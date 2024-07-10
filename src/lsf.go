package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readCsvFile(filePath string) ([][]string, error) {
    f, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    return csvReader.ReadAll()
}

func convertStringRecordsToFloat(records [][]string) ([][]float64, int, int, error) {
    var result [][]float64
    var err error
    var floatValue float64
    var n = len(records)
    var cols = len(records[0])
    for _, record := range records {
        var line []float64
        for _, value := range record {
            floatValue, err = strconv.ParseFloat(strings.TrimSpace(value), 64)
            if err != nil {
                return result, n, cols, err
            }
            line = append(line, floatValue)
        }

        result = append(result, line)
    }
    return result, n, cols, err
}

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

func leastSquaresFit(records [][]float64, xIndex int, yIndex int) (float64, float64, float64, float64, float64) {
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

func printData(data [][]float64) {
    for _, record := range data {
        fmt.Print("[")
        for _, value := range record {
            fmt.Print(value)
            fmt.Print("\t")
        }
        fmt.Println("]")
    }
    fmt.Println("")
}

var flagAlias = map[string]string{
    "v": "verbose",
}

func parseArgs()(string, *int, *int, *bool) {
    if len(os.Args) < 2 {
        fmt.Println("Usage: lsf -x <x column index> -y <y column index> -v <true/false> <csv file>")
        os.Exit(1)
    }
    
    xIndex := flag.Int("x", 0, "x column index")
    yIndex := flag.Int("y", 1, "y column index")
    verbose := flag.Bool("v", false, "verbose output")

    for from, to := range flagAlias {
        flagSet := flag.Lookup(from)
        flag.Var(flagSet.Value, to, fmt.Sprintf("alias to %s", flagSet.Name))
    }

    flag.Parse()
    fileName := (flag.Args())[0]

    return fileName, xIndex, yIndex, verbose
}

func readData(fileName string, xIndex *int, yIndex *int) (floatRecords [][]float64, n int) {
    records, err := readCsvFile(fileName)
    if err != nil {
        fmt.Println("Error reading CSV file", err)
        os.Exit(1)
    }

    floatRecords, n, cols, err := convertStringRecordsToFloat(records)
    if(err != nil) {
        fmt.Println("Error converting string records to float records", err)
        os.Exit(1)
    }

    if n < 3 {
        fmt.Println("Data set must have 3 or more points.")
        os.Exit(1)
    }

    if *xIndex < 0 || *xIndex >= cols {
        fmt.Println("x column index must be between 0 and", cols - 1, ".")
        os.Exit(1)
    }

    if *yIndex < 0 || *yIndex >= cols {
        fmt.Println("y column index must be between 0 and", cols - 1, ".")
        os.Exit(1)
    }

    return floatRecords, n
}

func maybeOutputData(verbose *bool, floatRecords [][]float64) {
    if *verbose {
      printData(floatRecords)
    }
}

func outputResult(m float64, dm float64, b float64, db float64, rr float64, n int) {
    fmt.Println("m:", m, "±", dm)
    fmt.Println("b:", b, "±", db)
    fmt.Println("rr:", rr)
    fmt.Println("n:", n)
}

func main() {
    fileName, xIndex, yIndex, verbose := parseArgs();

    fmt.Print("Least Squares Fit of ", fileName, " - column ", *yIndex, " versus column ", *xIndex, "\n")

    floatRecords, n := readData(fileName, xIndex, yIndex)
    m, dm, b, db, rSquared := leastSquaresFit(floatRecords, *xIndex, *yIndex)

    maybeOutputData(verbose, floatRecords)
    outputResult(m, dm, b, db, rSquared, n)
}