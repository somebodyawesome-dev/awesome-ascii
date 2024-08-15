package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime"

	. "github.com/somebodyawesome-dev/awesome-ascii.git/core"
	"github.com/somebodyawesome-dev/awesome-ascii.git/utils"
	"github.com/spf13/cobra"
)

var inputFile string
var width uint16
var asciiCharType AsciiCharType = Basic
var outputFile string
var concurrency int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "awesome-ascii",
	Short: "A image to ASCII CLI command",
	Long:  `A command to turn image into  ASCII texts.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		img, err := utils.OpenImage(inputFile)

		if err != nil {
			log.Fatalf("Error: %v", err)
			os.Exit(1)
		}
		scaledImage := ScaleImage(img, width)
		grayImage := ConvertToGrayscale(scaledImage)
		asciiArt := MapPixelsToASCII(grayImage, asciiCharType)

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

	rootCmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "", "An image path which will be converted to ASCII")
	rootCmd.MarkPersistentFlagRequired("input")

	termSize := utils.GetTerminalSize()

	rootCmd.PersistentFlags().Uint16VarP(&width, "width", "w", termSize.Col, "An image path which will be converted to ASCII")

	rootCmd.PersistentFlags().VarP(&asciiCharType, "ascii-type", "a", "Determine which set of ascii characters will be used")

	rootCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "", "An output path for the converted image")

	rootCmd.PersistentFlags().IntVarP(&concurrency, "concurrency", "c", runtime.NumCPU(), "Set GOMAXPROCS")
}
