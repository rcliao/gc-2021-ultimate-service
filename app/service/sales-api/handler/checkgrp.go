package handler

import (
	"context"
	"net/http"

	"github.com/rcliao/gc-2021-ultimate-service/foundation/web"
)

func readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "Ok",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}
