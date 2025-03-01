/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/bastienbyra/rolenv/internal/config"
	"github.com/bastienbyra/rolenv/internal/docker"
	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Checks if the configuration is valid",
	Run: func(cmd *cobra.Command, args []string) {
		dockerConfig, err := config.LoadConfig(cfgFile)
		if err != nil {
			log.Fatalf("Erreur lors du chargement de la config : %v", err)
		}
		docker.Validate(dockerConfig)
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
