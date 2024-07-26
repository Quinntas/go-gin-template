package createUser

import (
	"github.com/quinntas/go-gin-template/src/globals"
	userDomain "github.com/quinntas/go-gin-template/src/modules/core/user/domain"
	userRepo "github.com/quinntas/go-gin-template/src/modules/core/user/repo"
)

func UseCase(dto DTO) (ResponseDTO, error) {
	persistence := userDomain.User{
		Email:    dto.Email,
		Password: dto.Password,
		Role:     "CLIENT",
	}

	res, err := userRepo.InsertUser(globals.RWDatabase, persistence)
	if err != nil {
		return ResponseDTO{}, err
	}

	return NewResponseDTO(res.LastInsertId, "user created successfully"), nil
}
