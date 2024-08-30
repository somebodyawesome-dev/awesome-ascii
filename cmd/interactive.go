package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/somebodyawesome-dev/awesome-ascii.git/config"
	"github.com/somebodyawesome-dev/awesome-ascii.git/core"
	"github.com/somebodyawesome-dev/awesome-ascii.git/utils"
	"github.com/spf13/cobra"
)

func runSurvey() {
	var imageInput string
	var width string
	var asciiType string

	// Define the survey prompts
	imagePrompt := &survey.Input{
		Message: "Please enter the image path:",
	}
	widthPrompt := &survey.Input{
		Message: "Please enter the image width:",
	}

	keys := make([]string, 0, len(core.AsciiCharsMap))
	for key := range core.AsciiCharsMap {
		keys = append(keys, string(key))
	}

	asciiTypePrompt := &survey.Select{
		Message: "Choose an ASCII type:",
		Options: keys,
	}

	// Ask the user for input
	survey.AskOne(widthPrompt, &width)
	survey.AskOne(imagePrompt, &imageInput)
	survey.AskOne(asciiTypePrompt, &asciiType)

	// Convert the selected string back to the original type
	asciiTypeEnum := core.AsciiCharType(asciiType)

	processForm(imageInput, width, asciiTypeEnum)
}

func processForm(imageInput, width string, asciiType core.AsciiCharType) {
	img, err := utils.OpenImage(imageInput)

	if err != nil {
		log.Fatalf("Error: %v", err)
		os.Exit(1)
	}
	w, e := strconv.ParseUint(width, 10, 16)
	if e != nil {
		log.Fatalf("Error: %v", err)
		os.Exit(1)
	}

	newWidth := uint16(w)

	scaledImage := core.ScaleImage(img, newWidth)
	grayImage := core.ConvertToGrayscale(scaledImage)
	asciiArt := core.MapPixelsToASCII(grayImage, asciiType)

	if config.OutputFile != "" {
		utils.ToFile(asciiArt, config.OutputFile)
	} else {
		fmt.Println(asciiArt)
	}
}

var interactiveCMD = &cobra.Command{
	Use:   "interactive",
	Short: "Interactive mode",
	Long:  `Opens the CLI in interactive mode so that the user can choose the options manually.`,
	Run: func(cmd *cobra.Command, args []string) {
		runSurvey()
	},
}

func init() {
	rootCmd.AddCommand(interactiveCMD)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
