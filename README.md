# go-wasm-tiny

Minimal Go -> wasm running in browser experiment.

## Notes
wasm_exec should match your go version.
```bash
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" ./dist/
```

## Build
```bash
make
```
## Run
```bash
make run
# open http://localhost:8080
```
