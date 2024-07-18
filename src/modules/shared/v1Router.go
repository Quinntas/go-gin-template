package shared

import (
	"github.com/quinntas/go-gin-template/src/internal/api/web"
	"github.com/quinntas/go-gin-template/src/modules/user/infra/http"
)

func V1Router(app *web.App) {
	router := app.Engine.Group("/v1")

	http.Router(router, app)
}
