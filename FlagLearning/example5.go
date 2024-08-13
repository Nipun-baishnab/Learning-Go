package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var name string
	var age int

	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "A brief description of your application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello, %s!\n", name)
			if age > 0 {
				fmt.Printf("You are %d years old.\n".age)
			}
		},
	}

	rootCmd.Flags().StringVarP(&name, "name", "n", "World", "a name to say hello")
	rootCmd.Flags().IntVarP(&age, "age", "a", 0, "your age")

	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo a message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo:", args[0])
		},
	}

	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}
