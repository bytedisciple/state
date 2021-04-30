# Examples

## Compiling and Running

For each example folder run the following command to compile to web assembly.

`GOOS=js GOARCH=wasm go build .`

In the current working directory run the following command to serve the examples to your browser.

`go run main.go`

The examples should then be accessible in your browser on http://localhost:8080/.

## Simple

Simple is a basic example where two components depend on the same int state object. 