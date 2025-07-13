package dataReader

import (
	"testing"
)

func TestNonExistentCsvFile(t *testing.T) {
	_, err := Read("./non_existent_file.csv", 0, 1)
	if err.Error() != "error reading CSV file: open ./non_existent_file.csv: no such file or directory" {
		t.Fatalf("Expected specific error message, got: %v", err)
	}
}

func TestBadCsvData(t *testing.T) {
	_, err := Read("../data/notFloatData.csv", 0, 1)
	if err.Error() != "error converting string records to float records: strconv.ParseFloat: parsing \"a\": invalid syntax" {
		t.Fatalf("Expected specific error message, got: %v", err)
	}
}

func TestNotEnoughCsvData(t *testing.T) {
	_, err := Read("../data/notEnoughData.csv", 0, 1)
	if err.Error() != "data set must have 3 or more points" {
		t.Fatalf("Expected specific error message, got: %v", err)
	}
}

func TestXColBelowRange(t *testing.T) {
	_, err := Read("../data/data.csv", -1, 1)
	if err.Error() != "x column index must be between 0 and 2" {
		t.Fatalf("Expected specific error message, got: %v", err)
	}
}

func TestXColAboveRange(t *testing.T) {
	_, err := Read("../data/data.csv", 3, 1)
	if err.Error() != "x column index must be between 0 and 2" {
		t.Fatalf("Expected specific error message, got: %v", err)
	}
}

func TestYColBelowRange(t *testing.T) {
	_, err := Read("../data/data.csv", 0, -1)
	if err.Error() != "y column index must be between 0 and 2" {
		t.Fatalf("Expected specific error message, got: %v", err)
	}
}

func TestYColAboveRange(t *testing.T) {
	_, err := Read("../data/data.csv", 0, 3)
	if err.Error() != "y column index must be between 0 and 2" {
		t.Fatalf("Expected specific error message, got: %v", err)
	}
}

func TestHappyPath(t *testing.T) {
	_, err := Read("../data/data.csv", 0, 1)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
}