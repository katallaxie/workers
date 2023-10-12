package workers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"syscall/js"

	"github.com/katallaxie/workers/internal/utils"
)

var handler http.Handler

func init() {
	fn := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) > 2 {
			panic(fmt.Errorf("too many args given to handleRequest: %d", len(args)))
		}

		req := args[0]
		ctx := js.Null()
		if len(args) > 1 {
			ctx = args[1]
		}

		var cb js.Func
		cb = js.Func(func(_ js.Value, args []js.Value) any {
			defer cb.Release()

			go func() {
				res, err := handleRequest(req, ctx)
				if err != nil {
					panic(err)
				}
				resolve := args[0]
				resolve.Invoke(res)
			}()

			return utils.Undefined()
		})

		return utils.NewPromise(cb)
	})

	utils.Global.Set("handleRequest", fn)
}

func handleRequest(req, ctx js.Value) (js.Value, error) {
	if handleRequest == nil {
		return utils.Undefined(), fmt.Errorf("no handler registered")
	}

	req, err := utils.FromRequest(req)
	if err != nil {
		return utils.Undefined(), err
	}

	ctx = utils.WithRuntimeContext(context.Background(), ctx)
	req = req.WithContext(ctx)
	reader, writer := io.Pipe()

	go func() {
		handler.ServeHTTP()
	}

	return utils.Undefined()
}

// Serve serves the given handler.
func Serve(h http.Handler) (js.Value, error) {
	if handler == nil {
		handler = http.DefaultServeMux
	}

	handler = h
	utils.Global.Call("ready")

	select {}
}
