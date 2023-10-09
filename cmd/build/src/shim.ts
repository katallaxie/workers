import "./wasm_exec_go.js";

export const go = new Go();

let mod

export function init(m) {
  mod = m
}
