package userCreate

import (
	userDomain "github.com/quinntas/go-gin-template/src/modules/user/domain"
	userRepo "github.com/quinntas/go-gin-template/src/modules/user/repo"
)

func UseCase(dto DTO) error {
	user := userDomain.User{
		Email:    dto.Email,
		Password: dto.Password,
		Role:     "CLIENT",
	}

	_, err := userRepo.Create(user, dto.App)

	if err != nil {
		return err
	}

	return nil
}
