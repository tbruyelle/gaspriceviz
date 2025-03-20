build:
	go build .

wasm: build
	#go run gioui.org/cmd/gogio@main -target js -o wasm gaspriceviz
	gogio -target js -o wasm gaspriceviz

serve: wasm
	go run ./serve
