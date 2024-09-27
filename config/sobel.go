package config

import (
	"github.com/somebodyawesome-dev/awesome-ascii.git/core"
	"github.com/spf13/cobra"
)

func InitSobelConverter(cmd *cobra.Command) {

	InitBaseConverter(cmd)
	cmd.Flags().Uint8VarP(&core.SOBEL_THRESHOLD, "threshold", "t", core.SOBEL_THRESHOLD, "Threshold between 0..255 to control intensity of assci in the edges of the image")

}
