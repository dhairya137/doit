// cmd/list.go
package cmd

import (
	"doit/internal/installer"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available packages",
	Long:  `Display a list of all available packages that can be installed using doit`,
	Run: func(cmd *cobra.Command, args []string) {
		inst := installer.NewInstaller()
		packages := inst.ListAvailablePackages()

		// Group packages by category
		categoryMap := make(map[string][]installer.Package)
		for _, pkg := range packages {
			category := pkg.Category
			if category == "" {
				category = "Other"
			}
			categoryMap[category] = append(categoryMap[category], pkg)
		}

		// Get sorted categories
		categories := make([]string, 0, len(categoryMap))
		for category := range categoryMap {
			categories = append(categories, category)
		}
		sort.Strings(categories)

		// Create a new tabwriter
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)

		// Print each category and its packages
		for _, category := range categories {
			// Print category header
			fmt.Fprintf(w, "\n%s:\n", strings.ToUpper(category))
			fmt.Fprintln(w, strings.Repeat("-", len(category)+1))
			
			// Print packages in this category
			fmt.Fprintln(w, "PACKAGE\tDESCRIPTION")
			for _, pkg := range categoryMap[category] {
				fmt.Fprintf(w, "%s\t%s\n", pkg.Name, pkg.Description)
			}
		}

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}