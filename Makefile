.PHONY: build-local build templ notify-templ-proxy run

build-local:
	@go build -o ./tmp/main .

build:
	@go build -o ./tmp/main .

templ:
	@templ generate --watch --proxy=http://localhost:8080 --proxyport=8082 --open-browser=false --proxybind="0.0.0.0"

notify-templ-proxy:
	# https://templ.guide/developer-tools/live-reload-with-other-tools/#templ-watch-mode
	@templ generate --notify-proxy --proxyport=8082

run:
	@make templ & sleep 1
	@air