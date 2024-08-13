package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "myapp",
		Short: "My sample application using Cobra",
		Long:  "This is a sample application to demonstrate Cobra.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from myapp!")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
