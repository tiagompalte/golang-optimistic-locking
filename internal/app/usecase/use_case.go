package usecase

type UseCase struct {
	HealthCheckUseCase
	CreateUserUseCase
	ValidateUserPasswordUseCase
	GenerateUserTokenUseCase
	FindUserUUIDUseCase
}
