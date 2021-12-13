# cfgPath=/Users/adria/Programming/misou/appconfig.json
build:
	wails build -p -ldflags="-X 'github.com/elchead/misou/integration.cfgPath=${misouCfg}'"
index:
	go run ./scripts/indexer.go
