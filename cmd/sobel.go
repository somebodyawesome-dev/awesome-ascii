package cmd

import (
	"fmt"
	"os"

	"github.com/somebodyawesome-dev/awesome-ascii.git/config"
	"github.com/somebodyawesome-dev/awesome-ascii.git/core"
	"github.com/somebodyawesome-dev/awesome-ascii.git/utils"
	"github.com/spf13/cobra"
)

// sobelCmd represents the convert command
var sobelCmd = &cobra.Command{
	Use:   "sobel",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		img, err := utils.OpenImage(config.InputFile)

		if err != nil {
			fmt.Printf("rror: %v \n", err)
			os.Exit(1)
		}
		scaledImage := core.ScaleImage(img, config.Width)
		grayImage := core.ConvertToGrayscale(scaledImage)
		newImg := core.ApplySobel(grayImage)
		// asciiArt := ascii.MapPixelsToASCII(newImg.Gray, asciiCharType)

		asciiArt := newImg.ApplyEgdesToAscii()

		if config.OutputFile != "" {
			utils.ToFile(asciiArt, config.OutputFile)
		} else {
			fmt.Println(asciiArt)
		}

	},
}

func init() {
	rootCmd.AddCommand(sobelCmd)
	config.InitSobelConverter(sobelCmd)
}
