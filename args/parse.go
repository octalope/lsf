package args

import (
	"flag"
	"fmt"
	"os"
)

func Parse()(string, *int, *int, *bool) {
		var flagAlias = map[string]string{
				"v": "verbose",
		}

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
