package userMapper

import (
	"github.com/quinntas/go-gin-template/src/internal/ddd"
	"github.com/quinntas/go-gin-template/src/internal/utils"
	userDomain "github.com/quinntas/go-gin-template/src/modules/core/user/domain"
)

func ToDomain(raw utils.H) (userDomain.User, error) {
	id := raw["id"].(int)
	pid := raw["pid"].(string)
	email := raw["email"].(string)
	password := raw["password"].(string)
	role := raw["role"].(string)
	createdAtStr := raw["created_at"].(string)
	updatedAtStr := raw["updated_at"].(string)

	if id == 0 || pid == "" || email == "" || password == "" || role == "" || createdAtStr == "" || updatedAtStr == "" {
		return userDomain.User{}, ddd.InvalidEntity()
	}

	createdAt, err := utils.StringToTime(createdAtStr)
	if err != nil {
		return userDomain.User{}, err
	}

	updatedAt, err := utils.StringToTime(updatedAtStr)
	if err != nil {
		return userDomain.User{}, err
	}

	return userDomain.NewUser(
		ddd.NewBaseDomain(id, pid, createdAt, updatedAt),
		email,
		password,
		role,
	), nil
}
