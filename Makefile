build-cli:
	@go build -o ../myapp/capruk.exe ./cli

build-win:
	@go build -o ./dist/capruk.exe ./cli

build:
	@go build -o ./dist/capruk ./cli