package dataReader

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Read(fileName string, xIndex *int, yIndex *int) (floatRecords [][]float64, n int) {
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
