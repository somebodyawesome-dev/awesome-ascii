package core

import (
	"fmt"
	"image"
	"image/color"
	"strings"

	"github.com/somebodyawesome-dev/awesome-ascii.git/utils"
	"golang.org/x/image/draw"
)

type MapPixelParams struct {
	Colored    bool
	ColorImage image.Image
	Img        image.Gray
	AsciiType  AsciiCharType
	AsciiChar  rune
}

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

func MapPixelsToASCII(params MapPixelParams) string {
	bounds := params.Img.Bounds()
	asciiArt := make([][]string, bounds.Dy()) 
	asciiSet, err := params.AsciiType.GetAsciiChars()
	var asciiChar rune

	utils.ParallelImageProcess(bounds.Size(), func(x, y int) {
		grayColor := params.Img.GrayAt(x, y)
		colorPixel := params.ColorImage.At(x, y).(color.RGBA)
		if err != nil {
			asciiChar = params.AsciiChar
		} else {
			asciiChar = asciiSet[int(grayColor.Y)*len(asciiSet)/256]
		}
		// Allocate row if not yet done (thread-safe)
		if asciiArt[y] == nil {
			asciiArt[y] = make([]string, bounds.Dx())
		}
		if params.Colored {
			ansiColor := RGBToANSI(colorPixel.R, colorPixel.G, colorPixel.B)
			asciiArt[y][x] = fmt.Sprintf("%s%c%s", ansiColor, asciiChar, "\033[0m")
		} else {
			asciiArt[y][x] = string(asciiChar)
		}

	})

	result := ""
	for _, row := range asciiArt {
		result += strings.Join(row, "") + "\n"
	}
	return result
}

func RGBToANSI(r, g, b uint8) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}
