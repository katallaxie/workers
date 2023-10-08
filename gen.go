//go:build generate
// +build generate

package main

//go:generate cp ${GOROOT}/misc/wasm/wasm_exec.js ./cmd/gen/files/wasm_exec_go.js
//go:generate cp ${TINYGOROOT}/targets/wasm_exec.js ./cmd/gen/files/wasm_exec_tinygo.js
