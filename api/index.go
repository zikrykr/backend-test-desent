package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/zikrykr/backend-test-desent/app"
)

var Handler http.HandlerFunc

func init() {
	fiberApp := app.SetupApp()

	Handler = adaptor.FiberApp(fiberApp)
}
