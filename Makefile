# Antes de fazerr o deploy é necessário configurar o AWS profile
.PHONY: clean
clean:
	@go clean
	@rm -rf ./bin

# ldflag="-s -w" remove informações de debug do binário - vai deixar o binário mais leve
# https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications
.PHONY: build
build: clean
	env GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o bin/bootstrap

.PHONY: zip
zip: build
	@zip -j -9 bin/bootstrap.zip bin/bootstrap

.PHONY: format
format:
	@go fmt -s -w .