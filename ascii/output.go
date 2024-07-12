package ascii

import (
	"fmt"
	"log"
	"os"
)

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
