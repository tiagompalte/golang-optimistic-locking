package usecase

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewHealthCheckUseCaseImpl,
	NewCreateUserUseCaseImpl,
	NewValidateUserPasswordUseCaseImpl,
	NewGenerateUserTokenUseCaseImpl,
	NewFindUserUUIDUseCaseImpl,
)
