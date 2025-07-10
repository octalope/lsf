package dataReader

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Read(fileName string, xIndex int, yIndex int) (floatRecords [][]float64, e error) {
	records, err := readCsvFile(fileName)
	if err != nil {
		fmt.Println("Error reading CSV file", err)
		os.Exit(1)
	}

	floatRecords, n, cols, err := convertStringRecordsToFloat(records)

	if err != nil {
		return floatRecords, fmt.Errorf("error converting string records to float records: %w", err)
	}

	if n < 3 {
		return floatRecords, errors.New("data set must have 3 or more points")
	}

	if xIndex < 0 || xIndex >= cols {
		return floatRecords, fmt.Errorf("x column index must be between 0 and %d", cols-1)
	}

	if yIndex < 0 || yIndex >= cols {
		return floatRecords, fmt.Errorf("y column index must be between 0 and %d", cols-1)
	}

	return floatRecords, nil
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
