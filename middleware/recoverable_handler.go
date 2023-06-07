package middleware

import "net/http"

type Environment uint8

const (
	Prod Environment = iota
	Dev
)

type RecoverableHandler struct {
	*http.ServeMux
	Environment
}

type option func(*RecoverableHandler)

func NewRecoverableHandler(mux *http.ServeMux, opts ...option) *RecoverableHandler {
	rh := &RecoverableHandler{ServeMux: mux}
	for _, opt := range opts {
		opt(rh)
	}
	return rh
}

func DevEnv(rh *RecoverableHandler) {
	rh.Environment = Dev
}
