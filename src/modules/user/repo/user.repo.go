package userRepo

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/quinntas/go-gin-template/src/internal/api/web"
	userDomain "github.com/quinntas/go-gin-template/src/modules/user/domain"
)

func Create(user userDomain.User, app *web.App) (sql.Result, error) {
	q := fmt.Sprintf("INSERT INTO %s (pid, email, password, role) VALUES (?, ?, ?, ?)", "Users")
	return app.WDB.Exec(q, uuid.New().String(), user.Email, user.Password, user.Role)
}
