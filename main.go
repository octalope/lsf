package main

import (
	"fmt"

	"github.com/octalope/lsf/args"
	"github.com/octalope/lsf/dataReader"
	"github.com/octalope/lsf/stats"
)

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

func maybeOutputData(verbose bool, floatRecords [][]float64) {
	if verbose {
		printData(floatRecords)
	}
}

func outputResult(m float64, dm float64, b float64, db float64, rr float64, n int) {
	fmt.Printf("m: %.4f ± %.1g\n", m, dm)
	fmt.Printf("b: %.4f ± %.1g\n", b, db)
	fmt.Printf("rr: %.4f\n", rr)
	fmt.Printf("n: %d\n", n)
}

func main() {
	var args args.Arguments = args.Parse()

	fmt.Print("Least Squares Fit of ", args.FileName, " - column ", args.YIndex, " versus column ", args.XIndex, "\n")

	var floatRecords [][]float64
	var err error
	floatRecords, err = dataReader.Read(args.FileName, args.XIndex, args.YIndex)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	var m, dm, b, db, rSquared float64 = stats.LeastSquaresFit(floatRecords, args.XIndex, args.YIndex)

	maybeOutputData(args.Verbose, floatRecords)
	outputResult(m, dm, b, db, rSquared, len(floatRecords))
}
