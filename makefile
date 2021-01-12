default: build_server build_client

build_server:
	go build -o bin/daily-problem cmd/main.go

build_client:
	cd web && npm run server