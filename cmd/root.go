package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "doit",
	Short: "DevOps Installation Tools - A CLI tool for server setup",
	Long: `DevOps Installation Tools (doit) is a command-line utility that helps
DevOps engineers quickly set up and configure new servers with necessary packages
and tools. Currently supports Ubuntu with plans for more OS support.`,
}

func Execute() error {
	return rootCmd.Execute()
}
