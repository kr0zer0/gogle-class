package usecase

import (
	"gogle-class/backend/internal/usecase/usecase_contracts"
)

// all usecase interfaces here...

type UseCases struct {
	usecase_contracts.RegisterUseCase
	usecase_contracts.LoginUseCase
	usecase_contracts.RefreshTokensUseCase
}

func NewUseCases(registerUseCase usecase_contracts.RegisterUseCase, loginUseCase usecase_contracts.LoginUseCase, refreshTokensUseCase usecase_contracts.RefreshTokensUseCase) *UseCases {
	return &UseCases{
		RegisterUseCase:      registerUseCase,
		LoginUseCase:         loginUseCase,
		RefreshTokensUseCase: refreshTokensUseCase,
	}
}
