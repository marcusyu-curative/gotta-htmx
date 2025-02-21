.PHONY: build-local build templ notify-templ-proxy run

build-local:
	@go build -o ./tmp/main .

build:
	@npm run build
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ./tmp/main .

templ:
	@templ generate --watch --proxy=http://localhost:8080 --proxyport=8082 --open-browser=false --proxybind="0.0.0.0"

notify-templ-proxy:
	@templ generate --notify-proxy --proxyport=8082

run:
	@make templ & sleep 1
	@air