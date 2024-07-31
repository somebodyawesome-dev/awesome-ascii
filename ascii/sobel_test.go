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
	ApplySobel(img)
}
func BenchmarkSobelSeq(b *testing.B) {

	img, err := utils.OpenImage("../images/girl.jpg")
	if err != nil {
		b.Error(err)
	}
	ApplySobelSeq(img)

}
