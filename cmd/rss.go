/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/z9fr/ctftime-cli/pkg/handler"
	"github.com/z9fr/ctftime-cli/pkg/utils"
)

// rssCmd represents the rss command
var rssCmd = &cobra.Command{
	Use:   "rss",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			handler.DisplayUpcommingCTFRSS(5)
		} else {
			i, err := strconv.Atoi(args[0])
			if err != nil {
				utils.LogError(err)
			}
			handler.DisplayUpcommingCTFRSS(i)

		}

	},
}

func init() {
	rootCmd.AddCommand(rssCmd)
}
