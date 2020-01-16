package interfaces

// ImageService interface for ImageService
type ImageService interface {
	Open(path string)
	EmbedMessage(message []int)
	ExtractMessage() []int
}

// MessageService interface for MessageService
type MessageService interface {
	ReadMessage(path string) []int
	DecodeMessage(bytes []int, length int) string
}

// FileService interface for FileService
type FileService interface {
	FileExists(path string) bool
	BaseNoExt(path string) string
}

// ConversionService interface for ConversionService
type ConversionService interface {
	Uint8ToBinary(u uint8) []int
	BinaryToUint8(bin []int) uint8
	BinaryToChar(bin []int) string
}

// StegoService interface for StegoService
type StegoService interface {
	Hide(imageFile string, textFile string)
	Show(imageFile string, length int) string
}
