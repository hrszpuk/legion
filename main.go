package main

import (
	"flag"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
)

var characters = []string{"@", "#", "O", "o", "*", "°", ".", " "}
var info bool
var verbose bool

var verboseLog = func(message string) {
	if verbose {
		fmt.Println(message)
	}
}

func main() {
	// Check if an argument is given
	if len(os.Args) < 2 {
		fmt.Println("Usage: legion [flags] <path/to/image>")
		os.Exit(0)
	}

	flag.BoolVar(&info, "info", false, "shows additional file information on output")
	flag.BoolVar(&verbose, "v", false, "shows additional information about the image to ascii process")
	flag.Parse()

	verboseLog("- Command line flags parsed")

	verboseLog("- Fetching image data")
	data, format, err := getImageData(flag.Args()[0])
	if err != nil {
		fmt.Println("Error occurred when getting image data")
		os.Exit(0)
	}
	verboseLog("- Resizing and gray scaling image")
	img := grayscaleAndResize(data)
	verboseLog("- Rendering ascii image")
	render(img)
	if info {
		fmt.Println("Additional information:")
		showInfo(data, format)
	}
}

func showInfo(data image.Image, format string) {
	fmt.Printf("Filename: %s\nFormat: %s\nWidth: %d\nHeight: %d\n",
		flag.Args()[0], format, data.Bounds().Max.X, data.Bounds().Max.Y)
}

func grayscaleAndResize(data image.Image) *image.Gray {
	verboseLog("|  Resizing image")
	data = resize.Resize(64, 64, data, resize.Lanczos2)
	verboseLog("|  Creating gray scale image")
	img := image.NewGray(data.Bounds())
	verboseLog("|  Writing image data")
	for y := data.Bounds().Min.Y; y < data.Bounds().Max.Y; y++ {
		for x := data.Bounds().Min.X; x < data.Bounds().Max.X; x++ {
			img.Set(x, y, data.At(x, y))
		}
	}
	return img
}

// render as it's name may imply this function renders the ascii image to the command line
func render(img *image.Gray) {
	verboseLog("|  Getting image width and length")
	width := img.Bounds().Dx()
	length := img.Bounds().Dy()

	verboseLog("|  Writing image data to output\n")
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
	verboseLog("|  Opening path/to/file")
	f, err := os.Open(path)
	if err != nil {
		return nil, "unknown", err
	}
	defer f.Close()

	verboseLog("|  Decoding image format")
	data, format, err := image.Decode(f)
	if err != nil {
		return nil, "unknown", err
	}

	return data, format, err
}
