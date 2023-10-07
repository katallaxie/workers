package workers

import (
	"net/http"

	"github.com/katallaxie/workers/internal/utils"
)

var handler http.Handler

// Serve serves the given handler.
func Serve(h http.Handler) {
	if handler == nil {
		handler = http.DefaultServeMux
	}

	handler = h
	utils.Global.Call("ready")

	select {}
}
