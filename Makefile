.PHONY: dev dist

bun-install:
	@bun install --cwd frontend

dev: bun-install
	@mkdir -p tmp
	@bun run --cwd frontend dev --clearScreen false & air

dist: bun-install
	@mkdir -p tmp/dist
	@echo building frontend...
	@bun run --cwd frontend build --outDir ../backend/vite/dist --emptyOutDir
	@echo building application...
	@go build -o ./tmp/dist/filestorage -tags dist -ldflags="-s -w" ./backend/cmd/filestorage

