package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
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
			errorLog := fmt.Sprintf("%+v: %s", r, string(debug.Stack()))
			log.Println(errorLog)

			var respString strings.Builder
			respString.WriteString("Something went wrong\n")
			if rh.Environment == Dev {
				respString.WriteString(errorLog)
			}
			http.Error(w, respString.String(), http.StatusInternalServerError)
		}
	}()
	rh.ServeMux.ServeHTTP(w, r)
}
