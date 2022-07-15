/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/z9fr/ctftime-cli/pkg/handler"
)

// teamCmd represents the team command
var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a team ID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		handler.TeamDetails(args[0])
	},
}

func init() {
	rootCmd.AddCommand(teamCmd)
}
