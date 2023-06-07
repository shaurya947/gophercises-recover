package middleware

import (
	"log"
	"net/http"
)

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

func DevEnv(rh *RecoverableHandler) {
	rh.Environment = Dev
}

func NewRecoverableHandler(mux *http.ServeMux, opts ...option) *RecoverableHandler {
	rh := &RecoverableHandler{ServeMux: mux}
	for _, opt := range opts {
		opt(rh)
	}
	return rh
}

func (rh *RecoverableHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	rh.ServeMux.ServeHTTP(w, r)
}
