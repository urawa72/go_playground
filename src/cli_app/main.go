package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is root command

func main() {
	RootCmd := &cobra.Command{
		Use:   "culc",
		Short: "command line calculator",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("root command")
		},
    }

	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
}
