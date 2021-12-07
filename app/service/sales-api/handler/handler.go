// Package handlers manages the different versions of the API.
package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/dimfeld/httptreemux/v5"
	"go.uber.org/zap"
)

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(Shutdown chan os.Signal, Log *zap.SugaredLogger) *httptreemux.ContextMux {
	mux := httptreemux.NewContextMux()

	h := func(w http.ResponseWriter, r *http.Request) {
		status := struct {
			Status string
		}{
			Status: "Ok",
		}
		json.NewEncoder(w).Encode(status)
	}
	mux.Handle(http.MethodGet, "/test", h)

	return mux
}
