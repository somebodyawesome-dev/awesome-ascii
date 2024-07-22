package ascii

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
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

func mapPixelsToASCII(img image.Image, asciiType utils.AsciiCharType, ker Kernel) string {
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
			grayColor := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			asciiChar := asciiSet[int(grayColor.Y)*len(asciiSet)/256]
			asciiArt += string(asciiChar)
		}
		asciiArt += "\n"
	}

	return asciiArt
}

func ConvertImageToASCII(imagePath string, newWidth uint16, asciiType utils.AsciiCharType) (string, error) {
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
	asciiArt := mapPixelsToASCII(grayImage, asciiType)

	return asciiArt, nil
}

func ApplySobel(img image.Image) image.Gray {

	bounds := img.Bounds()

	dx := []int{-1, 0, 1, -1, 0, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 0, 1, 1, 1}
	gx := []int{-1, 0, 1, -2, 0, 2, -1, 0, 1}
	gy := []int{-1, -2, -1, 0, 0, 0, 1, 2, 1}

	resultImage := image.NewGray(bounds) // create new gray image to store convo results

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Max.X; x < bounds.Max.X; x++ {
			// skip edges and fill them with default value
			// this will help preserving original image size
			if x == bounds.Min.X || x == bounds.Max.X-1 || y == bounds.Min.Y || y == bounds.Max.Y-1 {
				resultImage.SetGray(x, y, color.Gray{Y: 0})
				continue
			}

		}
	}

}
