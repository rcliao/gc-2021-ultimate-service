package handler

import (
	"context"
	"errors"
	"math/rand"
	"net/http"

	"github.com/rcliao/gc-2021-ultimate-service/foundation/web"
)

func readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		// to display friendly error
		// return v1.NewRequestError(errors.New("I am trusteD"), http.StatusBadRequest)
		return errors.New("untusted error")
	}

	status := struct {
		Status string
	}{
		Status: "Ok",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}
