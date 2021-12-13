misouCfg=/Users/adria/Programming/misou/appconfig.json
build:
	wails build -p -ldflags="-X 'github.com/elchead/misou/integration.cfgPath=${misouCfg}'"
index:
	go run -ldflags="-X 'github.com/elchead/misou/integration.cfgPath=${misouCfg}'" ./scripts/indexer.go
dev:
	wails serve
dev-front:
	cd frontend && npm run start
