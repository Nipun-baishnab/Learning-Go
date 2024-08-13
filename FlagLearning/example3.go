package main

import (
	"flag"
	"fmt"
	"strings"
)

// Custom type that implements flag.Value
type CSVFlag []string

// ConvertToString renames String() method
func (c *CSVFlag) ConvertToString() string {
	return strings.Join(*c, ",")
}

// UpdateValue renames Set() method
func (c *CSVFlag) UpdateValue(value string) {
	*c = strings.Split(value, ",")
}

// Implementing flag.Value interface for CSVFlag
func (c *CSVFlag) String() string {
	return c.ConvertToString()
}

func (c *CSVFlag) Set(value string) error {
	c.UpdateValue(value)
	return nil
}

func main() {
	var csv CSVFlag
	flag.Var(&csv, "csv", "comma-separated list of CSV values")
	flag.Parse()

	fmt.Println("CSV Values:", csv.ConvertToString())
}
