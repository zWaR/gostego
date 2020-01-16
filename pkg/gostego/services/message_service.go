package services

import (
	"io/ioutil"
	"log"

	"github.com/zWaR/gostego/pkg/gostego/interfaces"
)

type message struct {
	fileService       interfaces.FileService
	conversionService interfaces.ConversionService
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetBits reads a text file.
func (message *message) ReadMessage(filepath string) []int {
	var bits []int
	if !message.fileService.FileExists(filepath) {
		return bits
	}
	data, err := ioutil.ReadFile(filepath)
	check(err)
	for _, c := range data {
		result := message.conversionService.Uint8ToBinary(c)
		bits = append(bits, result...)
	}
	return bits
}

// DecodeMessage decodes the given message.
func (message *message) DecodeMessage(bytes []int, length int) string {
	var bytesIndex = 0
	var word []int
	var text string
	for bytesIndex < len(bytes) {
		word = append(word, bytes[bytesIndex])
		bytesIndex++
		if len(word) == 8 {
			text += message.conversionService.BinaryToChar(word)
			word = nil
		}
	}
	if length == 0 {
		return text
	}
	return text[0:length]
}

// NewMessageService returns a new instance of the MessageService
func NewMessageService(fileService interfaces.FileService, conversionService interfaces.ConversionService) interfaces.MessageService {
	var messageInstance = new(message)
	messageInstance.fileService = fileService
	messageInstance.conversionService = conversionService
	var messageService interfaces.MessageService = messageInstance
	return messageService
}
