package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/somebodyawesome-dev/awesome-ascii.git/ascii"
	"github.com/somebodyawesome-dev/awesome-ascii.git/config"
	"github.com/somebodyawesome-dev/awesome-ascii.git/utils"
	"github.com/spf13/cobra"
)

var concurrency int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "awesome-ascii",
	Short: "A Image to ASCII application",
	Long:  `A Image to ASCII CLI application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		img, err := utils.OpenImage(config.InputFile)

		if err != nil {
			log.Fatalf("Error: %v", err)
			os.Exit(1)
		}
		scaledImage := ascii.ScaleImage(img, config.Width)
		grayImage := ascii.ConvertToGrayscale(scaledImage)
		asciiArt := ascii.MapPixelsToASCII(grayImage, config.AsciiCharType)

		if config.OutputFile != "" {
			utils.ToFile(asciiArt, config.OutputFile)
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
	rootCmd.PersistentFlags().IntVarP(&concurrency, "concurrency", "c", runtime.NumCPU(), "Set GOMAXPROCS")
	config.InitBaseConverter(rootCmd)
}
