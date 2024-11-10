package cmd

import (
	"fmt"
	"doit/internal/installer"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [package-name]",
	Short: "Install a package",
	Long:  `Install a specified package on the current system`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please specify a package to install")
			return
		}
		
		inst := installer.NewInstaller()
		if err := inst.Install(args[0]); err != nil {
			fmt.Printf("Failed to install package: %v\n", err)
			return
		}
		fmt.Printf("Successfully installed %s\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}