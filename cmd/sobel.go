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
	Short: "A command that applies the Sobel edge detection algorithm to images and converts the result into ASCII art.",
	Long: `part of a command-line interface (CLI) application designed to process images using the Sobel edge detection algorithm.
This command reads an input image, scales it to the desired width, converts it to grayscale, and then applies the Sobel algorithm to highlight the edges.
The resulting edge-detected image is then transformed into ASCII art.
Users can choose to display the ASCII art directly in the terminal or save it to a file.`,
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
