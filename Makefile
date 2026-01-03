DIST=./dist
WASM=$(DIST)/main.wasm

all: $(WASM)

$(WASM): main.go
	GOOS=js GOARCH=wasm go build -o $(WASM)

run: all
		cd $(DIST) && python3 -m http.server 8080

clean:
	rm -f $(WASM)
