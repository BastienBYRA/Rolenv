/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of the application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Rolenv version: 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
