package apifs

import "net/http"

type Definition struct {
	Description string
	Path        string // api path and fs path
	Handler     http.HandlerFunc
}
