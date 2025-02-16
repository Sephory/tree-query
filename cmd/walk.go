package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sephory/tree-query/tree"
	"github.com/spf13/cobra"
)

var walkCmd = &cobra.Command{
	Use:   "walk FILE_PATH",
	Short: "Walk and output the full AST of a file",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		language := tree.GetLanguageForFile(args[0])
		file, _ := os.ReadFile(args[0])
		t := tree.LoadTree(file, language)
		j, err := json.Marshal(t)
		cobra.CheckErr(err)
		fmt.Print(string(j))
	},
}

func init() {
	rootCmd.AddCommand(walkCmd)
}
