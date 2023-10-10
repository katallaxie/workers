/// <reference path='global.d.ts'/>

import "./wasm_exec_go";

export const go = new Go();

let mod

export function init(m) {
  mod = m
}
