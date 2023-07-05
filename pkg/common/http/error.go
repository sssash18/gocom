package http

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	AppCode        int64  `json:"code,omitempty"`
	ErrText        string `json:"error,omitempty "`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r,e.HTTPStatusCode)
	return nil
}

func ErrInternal(err error) render.Renderer {
	return &ErrResponse{
		Err : err,
		HTTPStatusCode: http.StatusInternalServerError,
		ErrText: err.Error(),

	}
}

func ErrBadRequest(err error) render.Renderer {
	return &ErrResponse{
		Err : err,
		HTTPStatusCode: http.StatusBadRequest,
		ErrText: err.Error(),
	}
}
