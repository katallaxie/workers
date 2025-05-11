package main

import (
	"fmt"
	"net/http"

	"github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/components/buttons"
	"github.com/katallaxie/workers"
)

func Demo() htmx.Node {
	return htmx.HTML5(
		htmx.HTML5Props{
			Title:    "Demo",
			Language: "en",
			Head:     []htmx.Node{},
		},
		htmx.Body(
			buttons.Button(buttons.ButtonProps{}, htmx.Text("Demo")),
		),
	)
}

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "%s", Demo())
	})

	workers.Serve(handler)
}
