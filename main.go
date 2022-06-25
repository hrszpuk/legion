package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
)

const MAX_INTENSITY = 3 * 65535

var characters = []string{"@", "#", "O", "o", "*", "°", ".", " "}

func main() {
	// Check if an argument is given
	if len(os.Args) < 2 {
		fmt.Println("Usage: legion path/to/image")
		os.Exit(0)
	}
	data, _, err := getImageData(os.Args[1])
	if err != nil {
		fmt.Println("Error occurred when getting image data")
		os.Exit(0)
	}
	img := grayscaleAndResize(data)
	render(img)
}

func showInfo()

func grayscaleAndResize(data image.Image) *image.Gray {
	data = resize.Resize(64, 64, data, resize.Lanczos2)
	img := image.NewGray(data.Bounds())
	for y := data.Bounds().Min.Y; y < data.Bounds().Max.Y; y++ {
		for x := data.Bounds().Min.X; x < data.Bounds().Max.X; x++ {
			img.Set(x, y, data.At(x, y))
		}
	}
	return img
}

// render as it's name may imply this function renders the ascii image to the command line
func render(img *image.Gray) {
	width := img.Bounds().Dx()
	length := img.Bounds().Dy()

	for y := 0; y < length; y++ {
		for x := 0; x < width; x++ {
			var gray float64 = float64(img.GrayAt(x, y).Y) / 255.0
			var index = math.Round(gray * float64(len(characters)-1))
			fmt.Print(string(characters[int(index)]))
		}
		fmt.Println()
	}

}

// getImageData reads the image and returns an image.Image
// this function is mainly to clean up code as this part has a lot of error handling
func getImageData(path string) (image.Image, string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, "unknown", err
	}
	defer f.Close()

	data, format, err := image.Decode(f)
	if err != nil {
		return nil, "unknown", err
	}

	return data, format, err
}
