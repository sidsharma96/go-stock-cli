/*
Copyright © 2024 Siddharth Sharma
*/
package cmd

import (
	"github.com/sidsharma96/go-stock-cli/api"
	"github.com/spf13/cobra"
)

var topgainersCmd = &cobra.Command{
	Use:   "topgainers",
	Short: "Reveals top 5 gainers in market",
	Long: `This command reveals top 5 gainers in stock market`,
	Run: func(cmd *cobra.Command, args []string) {
		api.GetTopGainers()
	},
}

func init() {
	rootCmd.AddCommand(topgainersCmd)
	topgainersCmd.Flags().BoolP("help", "h", false, "command is for getting top gainers in market")
}
