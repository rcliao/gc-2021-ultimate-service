// Package handler manages the different versions of the API.
package handler

import (
	"net/http"
	"os"

	"github.com/rcliao/gc-2021-ultimate-service/business/web/mid"
	"github.com/rcliao/gc-2021-ultimate-service/foundation/web"
	"go.uber.org/zap"
)

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(shutdown chan os.Signal, log *zap.SugaredLogger) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log))

	app.Handle(http.MethodGet, "/readiniess", readiness)

	return app
}
