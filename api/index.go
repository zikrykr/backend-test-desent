package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/zikrykr/backend-test-desent/app"
)

var fiberHandler http.HandlerFunc

func init() {
	fiberApp := app.SetupApp()

	fiberHandler = adaptor.FiberApp(fiberApp)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fiberHandler(w, r)
}
