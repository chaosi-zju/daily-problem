default: build_server build_client

build_server:
	go build -o output/daily-problem cmd/main.go
	cp conf/config_prod.yaml output/

build_client:
	cd web && npm install && npm run build