package args

import (
	"flag"
	"fmt"
	"os"
)

type Arguments struct {
    FileName string
    XIndex int
    YIndex int
    Verbose bool
}

func Parse() Arguments {
    var flagAlias = map[string]string{
        "v": "verbose",
    }

    if len(os.Args) < 2 {
        fmt.Println("Usage: lsf -x <x column index> -y <y column index> -v <true/false> <csv file>")
        os.Exit(1)
    }

    XIndex := flag.Int("x", 0, "x column index")
    YIndex := flag.Int("y", 1, "y column index")
    Verbose := flag.Bool("v", false, "verbose output")

    for from, to := range flagAlias {
        flagSet := flag.Lookup(from)
        flag.Var(flagSet.Value, to, fmt.Sprintf("alias to %s", flagSet.Name))
    }

    flag.Parse()
    FileName := (flag.Args())[0]

    return Arguments{FileName, *XIndex, *YIndex, *Verbose}
}
