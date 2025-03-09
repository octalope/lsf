# Linear Least Squares Fit

Welcome to the `lsf` repository!

## Description

This repository contains the source code for the `lsf` project. `lsf` is a command-line tool written in Go that calculates a Least-Squares Fit of data from a CSV file. 

## Installation

To install `lsf`, follow these steps:

1. Clone the repository: `git clone https://github.com/octalope/lsf.git`
2. Build the executable: `go build .`
3. Run `lsf` using the generated executable: `./lsf`

## Usage

To use `lsf`, run the following command:

```bash
./lsf [-xyv] [file]

  -verbose
        verbose output
  -x int
        x column index
  -y int
        y column index (default 1)
```

### Example

```bash
./lsf -x 0 -y 1 --verbose ./data/data.csv                
Least Squares Fit of ./data/data.csv - column 1 versus column 0
[1      1.03    5.09    ]
[2      1.97    8.04    ]
[3      3.01    10.79   ]
[4      3.94    14.14   ]
[5      4.97    16.88   ]
[6      6.02    19.93   ]

m: 0.9966 ± 0.0213
b: 0.0020 ± 0.0041
rr: 0.9999
n: 6
```

For more information on available options, run `lsf --help`.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
