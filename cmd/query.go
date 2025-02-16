/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sephory/tree-query/tree"
	"github.com/spf13/cobra"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		language := tree.GetLanguageForFile(args[1])
		file, err := os.ReadFile(args[1])
		cobra.CheckErr(err)
		matches, err := tree.QueryTree(file, language, args[0])
		cobra.CheckErr(err)
		maps := make([]map[string][]string, len(matches))
		for i, m := range matches {
			maps[i] = m.ToMap()
		}
		j, err := json.Marshal(maps)
		cobra.CheckErr(err)
		fmt.Println(string(j))
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
