//go:build !js

package workers

import (
	"fmt"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	Port string `envconfig:"PORT" default:"9900"`
}

// Serve is using a http.Handler as a handler for the Worker.
func Serve(handler http.Handler) {
	env := Env{}
	envconfig.MustProcess("", &env)

	if handler == nil {
		handler = http.DefaultServeMux
	}

	addr := fmt.Sprintf(":%s", env.Port)
	fmt.Printf("listening on: http://localhost%s\n", addr)

	http.ListenAndServe(addr, handler)
}

func ServeNonBlock(http.Handler) {
	panic("ServeNonBlock is not supported in non-JS environments")
}

func Ready() {
	panic("Ready is not supported in non-JS environments")
}

func Done() <-chan struct{} {
	panic("Done is not supported in non-JS environments")
}
