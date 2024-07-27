package cmd

import (
	"fmt"
	"image/jpeg"
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
		file, err := os.Open("./images/girl.jpg")
		if err != nil {
			fmt.Errorf("failed to open image: %v", err)
		}
		defer file.Close()

		img, err := jpeg.Decode(file)
		if err != nil {
			fmt.Errorf("failed to decode image: %v", err)
		}

		newImg := ascii.ApplySobel(img)	
		f, err := os.Create("sobel.jpg")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err = jpeg.Encode(f, &newImg, nil); err != nil {
			fmt.Printf("failed to encode: %v", err)
		}
		asciiResult,_ := ascii.ConvertImageToASCII("sobel.jpg",200,utils.Basic)

		fmt.Println(asciiResult)
	

	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
