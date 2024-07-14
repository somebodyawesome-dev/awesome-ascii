package cmd

import (
	"fmt"
	"log"
	"os"

	. "github.com/somebodyawesome-dev/awesome-ascii.git/ascii"
	"github.com/spf13/cobra"
)

var inputFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "awesome-ascii",
	Short: "A Image to ASCII application",
	Long:  `A Image to ASCII CLI application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		newWidth := 200
		asciiArt, err := ConvertImageToASCII(inputFile, newWidth)
		if err != nil {
			log.Fatalf("Error: %v", err)
			os.Exit(1)
		}

		fmt.Println(asciiArt)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.awesome-ascii.git.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "An image path which will be converted to ASCII")
	rootCmd.MarkFlagRequired("input")
}
