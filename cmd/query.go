package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sephory/tree-query/tree"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query QUERY_TEXT FILE_PATH",
	Short: "Query the AST of a file",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var out []byte
		text, err := cmd.Flags().GetBool("text")
		cobra.CheckErr(err)
		language := tree.GetLanguageForFile(args[1])
		file, err := os.ReadFile(args[1])
		cobra.CheckErr(err)
		matches, err := tree.QueryTree(file, language, args[0])
		cobra.CheckErr(err)
		if text {
			maps := make([]map[string]string, len(matches))
			for i, m := range matches {
				maps[i] = m.ToMap()
			}
			out, err = json.Marshal(maps)
		} else {
			out, err = json.Marshal(matches)
		}
		cobra.CheckErr(err)
		fmt.Print(string(out))
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
	queryCmd.PersistentFlags().BoolP("text", "t", false, "Just output text of captures")
}
