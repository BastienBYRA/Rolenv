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

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a container",
	Run: func(cmd *cobra.Command, args []string) {
		dockerConfig, err := config.LoadConfig(cfgFile)
		if err != nil {
			log.Fatalf("Erreur lors du chargement de la config : %v", err)
		}
		docker.Run(dockerConfig)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().StringP("config-file", "f", ".", "Path to the configuration file")
}
