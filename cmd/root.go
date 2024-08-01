package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime"

	. "github.com/somebodyawesome-dev/awesome-ascii.git/ascii"
	"github.com/somebodyawesome-dev/awesome-ascii.git/utils"
	"github.com/spf13/cobra"
)

var inputFile string
var width uint16
var asciiCharType utils.AsciiCharType = utils.Basic
var outputFile string
var concurrency int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "awesome-ascii",
	Short: "A Image to ASCII application",
	Long:  `A Image to ASCII CLI application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		img, err := utils.OpenImage(inputFile)

		if err != nil {
			log.Fatalf("Error: %v", err)
			os.Exit(1)
		}
		asciiArt := ConvertImageToASCII(img, width, asciiCharType)

		if outputFile != "" {
			utils.ToFile(asciiArt, outputFile)
		} else {
			fmt.Println(asciiArt)
		}
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//set concurrency 
		runtime.GOMAXPROCS(concurrency)
	},
	Version: "0.0.1-alpha",
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
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "An image path which will be converted to ASCII")
	rootCmd.MarkFlagRequired("input")

	termSize := utils.GetTerminalSize()

	rootCmd.Flags().Uint16VarP(&width, "width", "w", termSize.Col, "An image path which will be converted to ASCII")

	rootCmd.PersistentFlags().VarP(&asciiCharType, "ascii-type", "a", "Determine which set of ascii characters will be used")

	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "An output path for the converted image")

	rootCmd.PersistentFlags().IntVarP(&concurrency, "concurrency", "c", runtime.NumCPU(), "Set GOMAXPROCS")
}
