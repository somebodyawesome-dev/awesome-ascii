package config

import (
	"github.com/somebodyawesome-dev/awesome-ascii.git/core"
	"github.com/somebodyawesome-dev/awesome-ascii.git/utils"

	"github.com/spf13/cobra"
)

var InputFile string
var Width uint16
var AsciiCharType core.AsciiCharType = core.Basic
var OutputFile string

func InitBaseConverter(cmd *cobra.Command) {

	termSize := utils.GetTerminalSize()
	cmd.Flags().StringVarP(&InputFile, "input", "i", "", "An image path which will be converted to ASCII")
	cmd.MarkFlagRequired("input")
	cmd.Flags().Uint16VarP(&Width, "width", "w", termSize.Col, "An image path which will be converted to ASCII")
	cmd.Flags().VarP(&AsciiCharType, "ascii-type", "a", "Determine which set of ascii characters will be used")
	cmd.Flags().StringVarP(&OutputFile, "output", "o", "", "An output path for the converted image")

}
