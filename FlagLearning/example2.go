package main

import (
	"flag"
	"fmt"
	"strings"
)

// Custom type that implements flag.value
type CSVFlag []string // is a typedef

// String method for flag.Value interface
func (C *CSVFlag) String() string {
	return fmt.Sprint(*C)
}

/**/

// Set method for flag.Value interface
func (C *CSVFlag) Set(value string) error {
	*C = strings.Split(value, ",")
	return nil
}

func main() {
	var csv CSVFlag
	flag.Var(&csv, "csv", "comma-separated list of CSV values")
	flag.Parse()

	fmt.Println("CSV Values:", csv)
}
