package cmd

import (
	"fmt"
	"os"

	"github.com/somebodyawesome-dev/awesome-ascii.git/ascii"
	"github.com/somebodyawesome-dev/awesome-ascii.git/utils"
	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		img, err := utils.OpenImage(inputFile)

		if err != nil {
			fmt.Printf("rror: %v \n", err)
			os.Exit(1)
		}
		scaledImage := ascii.ScaleImage(img, width)
		grayImage := ascii.ConvertToGrayscale(scaledImage)
		newImg := ascii.ApplySobel(grayImage)
		asciiArt := ascii.MapPixelsToASCII(newImg.Gray, asciiCharType)

		asciiArt = newImg.ApplyEgdesToAscii(asciiArt)

		if outputFile != "" {
			utils.ToFile(asciiArt, outputFile)
		} else {
			fmt.Println(asciiArt)
		}

	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
}
