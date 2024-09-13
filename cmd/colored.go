package cmd

import (
	"fmt"
	"os"

	"github.com/somebodyawesome-dev/awesome-ascii.git/core"
	"github.com/somebodyawesome-dev/awesome-ascii.git/utils"
	"github.com/spf13/cobra"
)

var inputFile string
var width uint16
var char string

var coloredCMD = &cobra.Command{
	Use:   "colored",
	Short: "colored mode",
	Long:  `This mode generates the images in colored mode, in which the user only have to choose what character to replace pixels with`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(char) != 1 {
			fmt.Println("Error: Please provide exactly one character for the --char flag")
			os.Exit(1)
		}

		img, err := utils.OpenImage(inputFile)

		if err != nil {
			fmt.Printf("rror: %v \n", err)
			os.Exit(1)
		}
		scaledImage := core.ScaleImage(img, width)
		grayImage := core.ConvertToGrayscale(scaledImage)
		//NEED to find a solution for ascii char type (make it optional)
		asciiArt := core.MapPixelsToASCII(true, scaledImage, grayImage, nil)

		fmt.Println(asciiArt)
	},
}

func init() {
	rootCmd.AddCommand(coloredCMD)
	termSize := utils.GetTerminalSize()
	coloredCMD.Flags().StringVarP(&inputFile, "input", "i", "", "An image path which will be converted to ASCII")
	coloredCMD.MarkFlagRequired("input")
	coloredCMD.Flags().StringVarP(&char, "char", "H", "#", "An image path which will be converted to ASCII")
	coloredCMD.Flags().Uint16VarP(&width, "width", "w", termSize.Col, "An image path which will be converted to ASCII")
}
