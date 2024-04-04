/*
Copyright © 2024 Siddharth Sharma
*/
package cmd

import (
	"os"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/sidsharma96/go-stock-cli/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stock",
	Short: "CLI tool for stock market",
	Long: `A lightweight CLI tool to get stock price of a company you provide. Or get top gainers and losers in the market if required`,
	Args: cobra.MinimumNArgs(1),
	Example: "  stock google\n  stock topgainers",
	Run: func(cmd *cobra.Command, args []string) {
		api.GetSearchResult(args[0])
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cc.Init(&cc.Config{
        RootCmd:       rootCmd,
        Headings:      cc.HiCyan + cc.Bold + cc.Underline,
        Commands:      cc.HiGreen + cc.Bold,
        Example:       cc.Italic + cc.Yellow,
        ExecName:      cc.Bold,
        Flags:         cc.Bold,
		NoExtraNewlines: true,
    })

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}


