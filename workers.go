//go:build !js

package workers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kelseyhightower/envconfig"
)

const (
	defaultTimeout = 3 * time.Second
)

// Env is a struct that holds the environment variables for the server.
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

	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: defaultTimeout,
		Handler:           handler,
	}

	_ = server.ListenAndServe()
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
