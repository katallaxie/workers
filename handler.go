package workers

import (
	"net/http"

	"github.com/katallaxie/workers/internal/utils"
)

var handler http.Handler

// Serve serves the given handler.
func Serve(handler http.Handler) {
	if handler == nil {
		handler = http.DefaultServeMux
	}

	handler = handler
	utils.Global.Call("ready")

	select {}
}
