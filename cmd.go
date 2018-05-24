package mdflow

import (
	"github.com/spf13/cobra"
)

// ExecuteRoot executes the root command
func ExecuteRoot() error {
	return setupCommands().Execute()
}

func rootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mdflow",
		Short: "parse markdown files",
	}
}

func setupCommands() *cobra.Command {
	root := rootCmd()
	addParseCommand(root)
	return root
}

func addParseCommand(root *cobra.Command) {
	var uploadUrl string

	cmd := &cobra.Command{
		Use:   "parse [flags] [file]",
		Short: "parses a markdown file",
		RunE: func(c *cobra.Command, args []string) error {
			return ParseFile(args[0])
		},
		Args: cobra.MinimumNArgs(1),
	}

	cmd.Flags().StringVarP(&uploadUrl, "url", "u", "", "The confluence upload url for the converted markdown")
	root.AddCommand(cmd)
}
