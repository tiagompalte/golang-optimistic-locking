package server

import (
	"net/http"

	_ "github.com/tiagompalte/golang-clean-optimistic-locking/api"
	"github.com/tiagompalte/golang-clean-optimistic-locking/application"
	"github.com/tiagompalte/golang-clean-optimistic-locking/internal/pkg/server/routes"
)

// @title						Optimistic Locking API
// @version						1.0
// @description					Optimistic Locking API
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @contact.url					http://www.swagger.io/support
// @contact.email				support@swagger.io
// @license.name				Apache 2.0
// @license.url					http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath					/
// @schemes						http https
// @securityDefinitions.apikey 	BearerAuth
// @in 							header
// @name 						Authorization
func NewServer(app application.App) *http.Server {
	groupRoutes := routes.CreateRoute(app)
	return app.Server().NewServer(groupRoutes)
}
