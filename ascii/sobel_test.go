package ascii

import (
	"testing"

	"github.com/somebodyawesome-dev/awesome-ascii.git/utils"
)

func BenchmarkSobel(b *testing.B) {

	img, err := utils.OpenImage("../images/girl.jpg")
	if err != nil {
		b.Error(err)
	}
	grayImage := ConvertToGrayscale(img)
	ApplySobel(grayImage)
}
func BenchmarkSobelSeq(b *testing.B) {

	img, err := utils.OpenImage("../images/girl.jpg")
	if err != nil {
		b.Error(err)
	}
	grayImage := ConvertToGrayscale(img)
	ApplySobelSeq(grayImage)

}
