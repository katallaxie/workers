package workers

import (
	"fmt"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

// Env ...
type Env struct {
	Port string `envconfig:"PORT" default:"9900"`
}

// Serve is using a http.Handler as a handler for the Worker.
func Serve(handler http.Handler) {
	env := Env{}
	envconfig.Process("", &env)

	addr := fmt.Sprintf(":%s", env.Port)
	fmt.Printf("listening on: http://localhost%s\n", addr)

	http.ListenAndServe(addr, handler)
}
