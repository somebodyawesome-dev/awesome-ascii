package config

import (
	"github.com/somebodyawesome-dev/awesome-ascii.git/ascii"
	"github.com/spf13/cobra"
)

func InitSobelConverter(cmd *cobra.Command) {

	InitBaseConverter(cmd)
	cmd.Flags().Uint8VarP(&ascii.SOBEL_THRESHOLD, "threshold", "t", ascii.SOBEL_THRESHOLD, "Threshold between 0..255 to control intensity of assci in the edges of the image")

}
