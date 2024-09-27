package core

import (
	"errors"
	"image"
)

// Since AssciCharType is a new type, you can’t compare it to other types anymore
type AsciiCharType string

const (
	Basic      AsciiCharType = "basic"
	Binary                   = "binary"
	Contrast                 = "contrast"
	Extended                 = "extended"
	HighDetail               = "high_detail"
)

var validAsciiCharTypes = map[string]AsciiCharType{
	"basic":       Basic,
	"binary":      Binary,
	"contrast":    Contrast,
	"extended":    Extended,
	"high_detail": HighDetail,
}
var asciiCharsMap = map[AsciiCharType][]rune{
	Basic:      []rune(" .:-=+*#%@"),
	Binary:     []rune("10"),
	Contrast:   []rune(" ,.:%?S#@"),
	Extended:   []rune(" .:-~|*i!t2x6qZ%0B98W@"),
	HighDetail: []rune(" .`'^\",:;Il!i><~+_-?[]{}1()|/tjrfxnruvcxzXYUJCLQO0Zmwqpbkdhao*#MW&8%B$@"),
}

// String is used both by fmt.Print and by Cobra in help text
func (e *AsciiCharType) String() string {
	return string(*e)
}

// Set must have pointer receiver so it doesn't change the value of a copy
func (e *AsciiCharType) Set(v string) error {
	if val, ok := validAsciiCharTypes[v]; ok {
		*e = val
		return nil
	}

	return errors.New(`must be one of "basic", "binary", "contrast", "extended", or "high_detail"`)

}

// Type is only used in help text
func (e *AsciiCharType) Type() string {
	return "AsciiCharType"
}

func (e AsciiCharType) GetAsciiChars() ([]rune, error) {
	if chars, ok := asciiCharsMap[e]; ok {
		return chars, nil
	}
	return nil, errors.New("invalid AsciiCharType")
}

func ConvertImageToASCII(img image.Image, newWidth uint16, asciiType AsciiCharType) string {
	scaledImage := ScaleImage(img, newWidth)
	grayImage := ConvertToGrayscale(scaledImage)
	asciiArt := MapPixelsToASCII(grayImage, asciiType)
	return asciiArt
}
