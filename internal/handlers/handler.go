package handlers

import "github.com/julienschmidt/httprouter"

// Handler is an interface for models handlers
type Handler interface {
	Register(router *httprouter.Router)
}