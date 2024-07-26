package userRepo

import (
	"fmt"
	"github.com/quinntas/go-gin-template/src/internal/api/database"
	userDomain "github.com/quinntas/go-gin-template/src/modules/core/user/domain"
	userDatabase "github.com/quinntas/go-gin-template/src/modules/core/user/infra/database"
)

func InsertUser(db database.Database, user userDomain.User) (database.InsertResult, error) {
	return db.Insert(
		fmt.Sprintf("INSERT INTO %s (email, password, role) VALUES (?, ?, ?)", userDatabase.Table),
		user.Email,
		user.Password,
		user.Role,
	)
}
