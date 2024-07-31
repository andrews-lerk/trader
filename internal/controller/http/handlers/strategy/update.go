package strategy

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"io"
	"log/slog"
	"net/http"
	resp "trader/internal/controller/http/response"
)

type Request struct {
	Number int             `json:"number" validate:"required,min=1"`
	Params json.RawMessage `json:"params"`
}

type Response struct {
	resp.Response
	Msg string `json:"message"`
}

func Update(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.strategy.update.Update"

		logger := logger.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		logger.Info("Some")
		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			logger.Error("Empty body")
			render.JSON(w, r, resp.Error("Request is empty"))
			return
		}
		if err != nil {
			logger.Error("Json decode error")
			fmt.Println(err)
			render.JSON(w, r, resp.Error("Request decode failed"))
			return
		}
		render.JSON(w, r, Response{
			resp.OK(),
			"Update successfully add to broker queue",
		})
	}
}
