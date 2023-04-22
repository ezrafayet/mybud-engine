
# Copy the glue file to the assets folder
glue:
	cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" "./assets/wasm_exec.js"

# Build the wasm files
build:
	GOOS=js GOARCH=wasm go build -o ./assets/multiplier.wasm cmd/wasm/main.go

# Start the server that serves the web app
start:
	go run cmd/server/main.go
