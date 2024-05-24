include .env
export

dev:
	pnpm watch & gow -e=go,mod,html run .

default:
	pnpm build
	go run cmd/bundler/main.go