package services

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/zWaR/gostego/pkg/gostego/interfaces"
)

type data struct {
	fileService       interfaces.FileService
	conversionService interfaces.ConversionService
	imageData         image.Image
	imagePath         string
	pixels            []pixel
	newImage          *image.RGBA
}

type pixel struct {
	r []int
	g []int
	b []int
	a []int
}

// Open opens an image
func (data *data) Open(filepath string) {
	if !data.fileService.FileExists(filepath) {
		log.Fatal("Image does not exist!")
	}
	data.imagePath = filepath
	file, err := os.Open(filepath)
	check(err)
	defer file.Close()

	data.imageData, _, err = image.Decode(file)
	check(err)
	data.getPixels()
}

func (data *data) getPixels() {
	bounds := data.imageData.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := data.rgbaToPixel(data.imageData.At(x, y).RGBA())
			data.pixels = append(data.pixels, pixel)
		}
	}
}

// EmbedMessage embeds the given message to the pixels.
func (data *data) EmbedMessage(message []int) {
	var messageIndex = 0
	var pixelIndex = 0
	for messageIndex < len(message) {
		pixel := data.pixels[pixelIndex]
		if len(message) > messageIndex {
			pixel.r[len(pixel.r)-1] = message[messageIndex]
		}
		if len(message) > messageIndex+1 {
			pixel.g[len(pixel.g)-1] = message[messageIndex+1]
		}
		if len(message) > messageIndex+2 {
			pixel.b[len(pixel.b)-1] = message[messageIndex+2]
		}
		messageIndex += 3
		pixelIndex++
	}
	data.pixelToRgba()
	data.saveNewImage()
}

// ExtractMessage extract the message from the pixels.
func (data *data) ExtractMessage() []int {
	var bytes []int
	var pixelIndex = 0
	for pixelIndex < len(data.pixels) {
		pixel := data.pixels[pixelIndex]
		bytes = append(bytes, pixel.r[len(pixel.r)-1])
		bytes = append(bytes, pixel.g[len(pixel.g)-1])
		bytes = append(bytes, pixel.b[len(pixel.b)-1])
		pixelIndex++
	}
	return bytes
}

func (data *data) saveNewImage() {
	directory := filepath.Dir(data.imagePath)
	fileName := data.fileService.BaseNoExt(data.imagePath)
	newImageName := fileName + "_steg.png"
	newPath := filepath.Join(directory, newImageName)

	outputFile, err := os.Create(newPath)
	check(err)
	png.Encode(outputFile, data.newImage)
	outputFile.Close()
}

func (data *data) pixelToRgba() {
	data.newImage = image.NewRGBA(data.imageData.Bounds())
	var pixelIndex = 0
	bounds := data.imageData.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := data.pixels[pixelIndex]
			r := data.conversionService.BinaryToUint8(pixel.r)
			g := data.conversionService.BinaryToUint8(pixel.g)
			b := data.conversionService.BinaryToUint8(pixel.b)
			a := data.conversionService.BinaryToUint8(pixel.a)
			var c color.Color = color.RGBA{r, g, b, a}
			data.newImage.Set(x, y, c)
			pixelIndex++
		}
	}
}

func (data *data) rgbaToPixel(r uint32, g uint32, b uint32, a uint32) pixel {
	r8 := uint8(r / 257)
	g8 := uint8(g / 257)
	b8 := uint8(b / 257)
	a8 := uint8(a / 257)

	return pixel{
		data.conversionService.Uint8ToBinary(r8),
		data.conversionService.Uint8ToBinary(g8),
		data.conversionService.Uint8ToBinary(b8),
		data.conversionService.Uint8ToBinary(a8),
	}
}

// NewImageService is a ImageService provider
func NewImageService(fileService interfaces.FileService, conversionService interfaces.ConversionService) interfaces.ImageService {
	var dataInstance = new(data)
	dataInstance.fileService = fileService
	dataInstance.conversionService = conversionService
	var imageService interfaces.ImageService = dataInstance
	return imageService
}
