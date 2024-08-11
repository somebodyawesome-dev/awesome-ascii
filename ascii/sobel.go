package ascii

import (
	"image"
	"image/color"
	"math"
	"strings"

	"github.com/somebodyawesome-dev/awesome-ascii.git/utils"
)

type SobelImage struct {
	image.Gray
	edgesAngle [][]float64
}

func (s SobelImage) GetEdgesAngleAt(x, y int) float64 {
	return s.edgesAngle[y][x]
}

func (s SobelImage) ApplyEgdesToAscii(asciiArt string) string {

	rows := strings.Split(asciiArt, "\n")

	for y := 0; y < len(rows); y++ {
		line := []rune(rows[y])

		for x := 0; x < len(rows[y]); x++ {

			edge := s.GetEdgesAngleAt(x, y)

			if (edge >= 0 && edge < 0.05) || (edge > 0.9 && edge <= 1) {

				line[x] = '_'
			} else if edge >= 0.45 && edge <= 0.55 {
				line[x] = '|'
			} else if edge >= 0.25 && edge < 0.3 {
				line[x] = '\\'
			} else if edge >= 0.75 && edge <= 0.9 {
				line[x] = '/'
			}

		}
		rows[y] = string(line)

	}
	return strings.Join(rows, "\n")

}

func ApplySobel(img image.Gray) SobelImage {

	bounds := img.Bounds()

	dx := []int{-1, 0, 1, -1, 0, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 0, 1, 1, 1}
	gx := []int{-1, 0, 1, -2, 0, 2, -1, 0, 1}
	gy := []int{-1, -2, -1, 0, 0, 0, 1, 2, 1}

	resultImage := image.NewGray(bounds) // create new gray image to store convo results

	magnitudes := make([][]float64, bounds.Max.Y)
	angles := make([][]float64, bounds.Max.Y)
	maxMagnitude := 0.0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		//init slices
		magnitudes[y] = make([]float64, bounds.Max.X)
		angles[y] = make([]float64, bounds.Max.X)

	}
	utils.ParallelImageProcess(bounds.Size(), func(x, y int) {

		// skip edges and fill them with default value
		// this will help preserving original image size
		if x == bounds.Min.X || x == bounds.Max.X-1 || y == bounds.Min.Y || y == bounds.Max.Y-1 {
			resultImage.SetGray(x, y, color.Gray{Y: 0})
			return
		}

		var sobelX, sobelY int
		for kernelIndex := 0; kernelIndex < 9; kernelIndex++ {

			pixelXIndex := x + dx[kernelIndex]
			pixelYIndex := y + dy[kernelIndex]

			grayValue := img.GrayAt(pixelXIndex, pixelYIndex).Y

			sobelX += gx[kernelIndex] * int(grayValue)
			sobelY += gy[kernelIndex] * int(grayValue)

		}
		magnitudes[y][x] = math.Sqrt(float64(sobelX*sobelX + sobelY*sobelY))
		angles[y][x] = 0.5*math.Atan2(float64(sobelY), float64(sobelX))/math.Pi + 0.5 // result in range [0, 1]

		if magnitudes[y][x] > maxMagnitude {
			maxMagnitude = magnitudes[y][x]
		}

	})

	utils.ParallelImageProcess(bounds.Size(), func(x, y int) {

		// normilize magnitude valuese
		magnitude := uint8((1 - magnitudes[y][x]/maxMagnitude) * 255)
		resultImage.SetGray(x, y, color.Gray{Y: magnitude})
	})

	return SobelImage{Gray: *resultImage, edgesAngle: angles}

}

func ApplySobelSeq(img image.Gray) SobelImage {

	bounds := img.Bounds()

	dx := []int{-1, 0, 1, -1, 0, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 0, 1, 1, 1}
	gx := []int{-1, 0, 1, -2, 0, 2, -1, 0, 1}
	gy := []int{-1, -2, -1, 0, 0, 0, 1, 2, 1}

	resultImage := image.NewGray(bounds) // create new gray image to store convo results

	magnitudes := make([][]float64, bounds.Max.Y)
	angles := make([][]float64, bounds.Max.Y)
	maxMagnitude := 0.0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		magnitudes[y] = make([]float64, bounds.Max.X)
		angles[y] = make([]float64, bounds.Max.X)
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// skip edges and fill them with default value
			// this will help preserving original image size
			if x == bounds.Min.X || x == bounds.Max.X-1 || y == bounds.Min.Y || y == bounds.Max.Y-1 {
				resultImage.SetGray(x, y, color.Gray{Y: 0})
				continue
			}

			var sobelX, sobelY int
			for kernelIndex := 0; kernelIndex < 9; kernelIndex++ {

				pixelXIndex := x + dx[kernelIndex]
				pixelYIndex := y + dy[kernelIndex]

				grayValue := img.GrayAt(pixelXIndex, pixelYIndex).Y

				sobelX += gx[kernelIndex] * int(grayValue)
				sobelY += gy[kernelIndex] * int(grayValue)

			}
			magnitudes[y][x] = math.Sqrt(float64(sobelX*sobelX + sobelY*sobelY))

			angles[y][x] = 0.5*math.Atan2(float64(sobelY), float64(sobelX))/math.Pi + 0.5 // result in range [0, 1]

			if magnitudes[y][x] > maxMagnitude {
				maxMagnitude = magnitudes[y][x]
			}

		}
	}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			// normilize magnitude valuese
			magnitude := uint8((1 - magnitudes[y][x]/maxMagnitude) * 255)
			resultImage.SetGray(x, y, color.Gray{Y: magnitude})
		}
	}

	return SobelImage{Gray: *resultImage, edgesAngle: angles}

}
