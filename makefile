default: build_server build_client

build_server:
	go build -o output/main cmd/main.go
	cp conf/config_prod.yaml output/

build_client:
	cd web && npm install && npm run build