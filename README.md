# WASM with Go

- Be sure having Go installed and the GOROOT path is exist
- Copy JS glue by `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .`
- `GOOS=js GOARCH=wasm go build -o main.wasm`

Note: If VSCode complains about "syscall/js" package, [configure as in here](https://gist.github.com/mdtanrikulu/c90cfdfe7bc7ccd8228e22f8586fb83c)

## Serve html + wasm in local with dev http-server
- Run web server `go run cmd/server.go`
- Visit `http://localhost:7532`
