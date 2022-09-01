//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializedServices() *simpleServices {
	wire.Build(
		newSimpleRepository, newSimpleServices,
	)
	return nil
}
