default: build_server build_client

build_server:
	go build -o output/server/main cmd/main.go
	cp conf/config_prod.yaml output/server/
	cp conf/daily_problem.service output/server/

build_client:
	cd web && npm install && npm run build