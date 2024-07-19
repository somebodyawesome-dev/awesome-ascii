package ascii

import (
	"fmt"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

var asciiChars = []rune("@%#*+=-:. ")

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

func convertToGrayscale(img image.Image) image.Image {
	bounds := img.Bounds()
	grayImage := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayColor := color.GrayModel.Convert(img.At(x, y))
			grayImage.Set(x, y, grayColor)
		}
	}

	return grayImage
}

func mapPixelsToASCII(img image.Image) string {
	bounds := img.Bounds()
	// width := bounds.Dx()
	asciiArt := ""

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayColor := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			asciiChar := asciiChars[int(grayColor.Y) * len(asciiChars) / 256]
			asciiArt += string(asciiChar)
		}
		asciiArt += "\n"
	}

	return asciiArt
}

func ConvertImageToASCII(imagePath string, newWidth uint16) (string, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return "", fmt.Errorf("failed to open image: %v", err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %v", err)
	}

	scaledImage := scaleImage(img, newWidth)
	grayImage := convertToGrayscale(scaledImage)
	asciiArt := mapPixelsToASCII(grayImage)

	return asciiArt, nil
}
