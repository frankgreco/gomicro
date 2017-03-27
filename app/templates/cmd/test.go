package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

func init() {
    RootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
    Use:   `test`,
    Short: `run all unit tests`,
    Long:  `run all unit tests`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("operation not yet implemented")
    },
}
