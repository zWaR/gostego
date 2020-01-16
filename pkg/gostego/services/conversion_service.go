package services

import (
	"fmt"
	"strconv"

	"github.com/zWaR/gostego/pkg/gostego/interfaces"
)

type conversionData struct{}

// Uint8ToBinary converts an uint8 integer to an array of binary integers.
func (conversionData *conversionData) Uint8ToBinary(u uint8) []int {
	var bin string
	bin = fmt.Sprintf("%s%.8b", bin, u)
	var bits []int
	for _, r := range bin {
		char := string(r)
		integer, _ := strconv.Atoi(char)
		bits = append(bits, integer)
	}
	return bits
}

// BinaryToUint8 converts binary values from an array to uint8.
func (conversionData *conversionData) BinaryToUint8(bin []int) uint8 {
	var index = 0
	var bit string
	var byteSequence string
	for index < 8 {
		bit = fmt.Sprintf("%d", bin[index])
		byteSequence += bit
		index++
	}
	integer, _ := strconv.ParseInt(byteSequence, 2, 64)
	return uint8(integer)
}

// BinaryToChar converts a sequence of 8 bytes to a ASCII char.
// Returns only characters between ASCII codes 31 and 123.
func (conversionData *conversionData) BinaryToChar(bin []int) string {
	charCode := conversionData.BinaryToUint8(bin)
	if charCode > 31 && charCode < 123 {
		return string(charCode)
	}
	return ""
}

// NewConversionService is a ConversionService providers
func NewConversionService() interfaces.ConversionService {
	var conversionDataInstance = new(conversionData)
	var conversionService interfaces.ConversionService = conversionDataInstance
	return conversionService
}
