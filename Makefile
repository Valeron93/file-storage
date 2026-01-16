.PHONY: dev dist bun-install

dev: bun-install
	@mkdir -p tmp
	@bun run --cwd frontend dev --clearScreen false & air

bun-install:
	@bun install --cwd frontend

dist: bun-install
	@mkdir -p tmp/dist
	@echo building frontend...
	@bun run --cwd frontend build --emptyOutDir
	@echo building application...
	@go build -o ./tmp/dist/filestorage -tags dist -ldflags="-s -w" ./cmd/filestorage

