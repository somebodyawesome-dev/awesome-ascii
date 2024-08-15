package cmd

import (
	"fmt"
	"os"

	"github.com/somebodyawesome-dev/awesome-ascii.git/core"
	"github.com/somebodyawesome-dev/awesome-ascii.git/utils"
	"github.com/spf13/cobra"
)

// sobelCmd represents the convert command
var sobelCmd = &cobra.Command{
	Use:   "sobel",
	Short: "Image to ASCII Text using Sobel filter",
	Long:  `A CLI command to turn images to ASCII texts by grayscaling and  applying Sobel filter to the input image.`,
	Run: func(cmd *cobra.Command, args []string) {

		img, err := utils.OpenImage(inputFile)

		if err != nil {
			fmt.Printf("rror: %v \n", err)
			os.Exit(1)
		}
		scaledImage := core.ScaleImage(img, width)
		grayImage := core.ConvertToGrayscale(scaledImage)
		newImg := core.ApplySobel(grayImage)
		// asciiArt := ascii.MapPixelsToASCII(newImg.Gray, asciiCharType)

		asciiArt := newImg.ApplyEgdesToAscii()

		if outputFile != "" {
			utils.ToFile(asciiArt, outputFile)
		} else {
			utils.ToStd(asciiArt)
		}

	},
}

func init() {
	rootCmd.AddCommand(sobelCmd)
	sobelCmd.Flags().Uint8VarP(&core.SOBEL_THRESHOLD, "threshold", "t", core.SOBEL_THRESHOLD, "Threshold between 0..255 to control intensity of assci in the edges of the image")
}
