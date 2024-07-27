package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func OpenImage(filepath string) (image.Image, error) {

	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image: %v", err)
	}
	defer file.Close()

	_, format, err := image.DecodeConfig(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image config: %v", err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to seek file: %v", err)
	}

	var img image.Image
	switch format {
	case "jpeg":
		img, err = jpeg.Decode(file)
		if err != nil {
			return nil, fmt.Errorf("failed to decode JPEG image: %v", err)
		}
	case "png":
		img, err = png.Decode(file)
		if err != nil {
			return nil, fmt.Errorf("failed to decode PNG image: %v", err)
		}
	default:
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}
	return img, nil

}

func ToStd(ascii string) (int, error) {
	return fmt.Println(ascii)
}

func ToFile(ascii string, outputFile string) {

	output := []byte(ascii)

	err := os.WriteFile(outputFile, output, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
