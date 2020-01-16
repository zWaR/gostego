//+build wireinject

package providers

import (
	"github.com/google/wire"
	"github.com/zWaR/gostego/pkg/gostego/interfaces"
	"github.com/zWaR/gostego/pkg/gostego/services"
)

// CreateFileService is a FileService provider.
func CreateFileService() interfaces.FileService {
	panic(
		wire.Build(
			services.NewFileService,
		),
	)
}

// CreateMessageService is a MessageService provider.
func CreateMessageService() interfaces.MessageService {
	panic(
		wire.Build(
			services.NewFileService,
			services.NewConversionService,
			services.NewMessageService,
		),
	)
}

// CreateImageService is a ImageService provider.
func CreateImageService() interfaces.ImageService {
	panic(
		wire.Build(
			services.NewFileService,
			services.NewConversionService,
			services.NewImageService,
		),
	)
}

// CreateConversionService is a ConversionService provider.
func CreateConversionService() interfaces.ConversionService {
	panic(
		wire.Build(
			services.NewConversionService,
		),
	)
}

// CreateStegoService is a StegoService provider.
func CreateStegoService() interfaces.StegoService {
	panic(
		wire.Build(
			services.NewFileService,
			services.NewConversionService,
			services.NewMessageService,
			services.NewImageService,
			services.NewStegoService,
		),
	)
}
