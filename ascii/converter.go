package ascii

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/somebodyawesome-dev/awesome-ascii.git/utils"
	"golang.org/x/image/draw"
)

func scaleImage(img image.Image, newWidth uint16) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	aspectRatio := float64(height) / float64(width)
	newHeight := int(aspectRatio * float64(newWidth))

	scaledImage := image.NewRGBA(image.Rect(0, 0, int(newWidth), newHeight))
	draw.CatmullRom.Scale(scaledImage, scaledImage.Bounds(), img, bounds, draw.Over, nil)

	return scaledImage
}

func convertToGrayscale(img image.Image) image.Gray {
	bounds := img.Bounds()
	grayImage := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayColor := color.GrayModel.Convert(img.At(x, y))
			grayImage.Set(x, y, grayColor)
		}
	}

	return *grayImage
}

func mapPixelsToASCII(img image.Gray, asciiType utils.AsciiCharType) string {
	bounds := img.Bounds()
	// width := bounds.Dx()
	asciiArt := ""
	asciiSet, err := asciiType.GetAsciiChars()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayColor := img.GrayAt(x, y)
			asciiChar := asciiSet[int(grayColor.Y)*len(asciiSet)/256]
			asciiArt += string(asciiChar)
		}
		asciiArt += "\n"
	}

	return asciiArt
}

func ConvertImageToASCII(img image.Image, newWidth uint16, asciiType utils.AsciiCharType) string {
	scaledImage := scaleImage(img, newWidth)
	grayImage := convertToGrayscale(scaledImage)
	asciiArt := mapPixelsToASCII(grayImage, asciiType)
	return asciiArt
}
