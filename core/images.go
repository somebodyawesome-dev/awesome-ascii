package core

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/somebodyawesome-dev/awesome-ascii.git/utils"
	"golang.org/x/image/draw"
)

func ScaleImage(img image.Image, newWidth uint16) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	aspectRatio := float64(height) / float64(width)
	newHeight := int(aspectRatio * float64(newWidth))

	scaledImage := image.NewRGBA(image.Rect(0, 0, int(newWidth), newHeight))
	draw.CatmullRom.Scale(scaledImage, scaledImage.Bounds(), img, bounds, draw.Over, nil)

	return scaledImage
}

func ConvertToGrayscale(img image.Image) image.Gray {
	bounds := img.Bounds()
	grayImage := image.NewGray(bounds)

	utils.ParallelImageProcess(bounds.Size(), func(x, y int) {
		grayColor := color.GrayModel.Convert(img.At(x, y))
		grayImage.Set(x, y, grayColor)
	})

	return *grayImage
}

func MapPixelsToASCII(colored bool, colorImage image.Image, img image.Gray, asciiType AsciiCharType) string {
	bounds := img.Bounds()
	asciiArt := ""
	asciiSet, err := asciiType.GetAsciiChars()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// TODO: Find work around to use parellel processing
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayColor := img.GrayAt(x, y)
			colorPixel := colorImage.At(x, y).(color.RGBA)
			asciiChar := asciiSet[int(grayColor.Y)*len(asciiSet)/256]

			if colored {
				ansiColor := RGBToANSI(colorPixel.R, colorPixel.G, colorPixel.B)
				asciiArt += fmt.Sprintf("%s%c%s", ansiColor, asciiChar, "\033[0m")
			} else {
				asciiArt += string(asciiChar)
			}
		}
		asciiArt += "\n"
	}

	return asciiArt
}

func RGBToANSI(r, g, b uint8) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func GrayToASCII(gray uint8, asciiType AsciiCharType) rune {
	return '#'
}
