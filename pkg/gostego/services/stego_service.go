package services

import "github.com/zWaR/gostego/pkg/gostego/interfaces"

type stego struct {
	imageService   interfaces.ImageService
	messageService interfaces.MessageService
}

// Hide hides the message in the image.
func (stego *stego) Hide(imageFile string, textFile string) {
	stego.imageService.Open(imageFile)
	message := stego.messageService.ReadMessage(textFile)
	stego.imageService.EmbedMessage(message)
}

// Show extracts the message from the image.
func (stego *stego) Show(imageFile string, length int) string {
	stego.imageService.Open(imageFile)
	message := stego.imageService.ExtractMessage()
	text := stego.messageService.DecodeMessage(message, length)
	return text
}

// NewStegoService is the stego service provider.
func NewStegoService(imageService interfaces.ImageService, messageService interfaces.MessageService) interfaces.StegoService {
	var stegoInstance = new(stego)
	stegoInstance.imageService = imageService
	stegoInstance.messageService = messageService
	var stegoService interfaces.StegoService = stegoInstance
	return stegoService
}
