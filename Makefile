include .env
export

# for hot-reloading install https://github.com/mitranim/gow
dev:
	pnpm watch & gow -e=go,mod,html run .

default:
	pnpm build
	go build -o example main.go