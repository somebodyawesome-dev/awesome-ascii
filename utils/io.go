package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)


// OpenImage opens and decodes an image file from the given filepath.
func OpenImage(filepath string) (image.Image, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image: %v", err)
	}
	defer file.Close()

	// Check the image format.
	_, format, err := image.DecodeConfig(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image config: %v", err)
	}

	// Reset the file pointer to the beginning of the file.
	if _, err := file.Seek(0, 0); err != nil {
		return nil, fmt.Errorf("failed to seek file: %v", err)
	}

	// Decode the image based on its format.
	switch format {
	case "jpeg":
		img, err := jpeg.Decode(file)
		if err != nil {
			return nil, fmt.Errorf("failed to decode JPEG image: %v", err)
		}
		return img, nil
	case "png":
		img, err := png.Decode(file)
		if err != nil {
			return nil, fmt.Errorf("failed to decode PNG image: %v", err)
		}
		return img, nil
	default:
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}
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
