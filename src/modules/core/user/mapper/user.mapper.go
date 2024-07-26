package userMapper

import (
	"github.com/quinntas/go-gin-template/src/internal/utils"
	userDomain "github.com/quinntas/go-gin-template/src/modules/core/user/domain"
)

func ToDomain(raw utils.H) (userDomain.User, error) {
	return userDomain.User{}, nil
}
