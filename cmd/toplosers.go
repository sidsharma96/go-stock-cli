/*
Copyright © 2024 Siddharth Sharma
*/
package cmd

import (
	"github.com/sidsharma96/go-stock-cli/api"
	"github.com/spf13/cobra"
)

var toplosersCmd = &cobra.Command{
	Use:   "toplosers",
	Short: "Reveals top 5 losers in market",
	Long: `This command reveals top 5 losers in stock market`,
	Run: func(cmd *cobra.Command, args []string) {
		api.GetTopLosers()
	},
}

func init() {
	rootCmd.AddCommand(toplosersCmd)
	toplosersCmd.Flags().BoolP("help", "h", false, "command is for getting top losers in market")
}
