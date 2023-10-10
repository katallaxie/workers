import * as shim from "./shim"
import mod from "./app.wasm"

shim.init(mod)

export default {}
