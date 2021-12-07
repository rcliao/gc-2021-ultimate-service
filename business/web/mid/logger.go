package mid

import (
	"context"
	"net/http"

	"github.com/rcliao/gc-2021-ultimate-service/foundation/web"
	"go.uber.org/zap"
)

func Logger(log *zap.SugaredLogger) web.Middleware {

	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			// LOG STARTED
			log.Infow("request started", "method", r.Method, "path", r.URL.Path,
				"remoteaddr", r.RemoteAddr)

			err := handler(ctx, w, r)

			// LOG COMPLETED
			log.Infow("request completed", "method", r.Method, "path", r.URL.Path,
				"remoteaddr", r.RemoteAddr)

			return err
		}

		return h
	}

	return m
}
