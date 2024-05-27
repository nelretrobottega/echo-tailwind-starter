# Echo-tailwind-starter

Starter project for integrating tailwind (daisy UI) with golang templates using Echo framework.

![image](https://github.com/nelretrobottega/echo-tailwind-starter/assets/35533749/404d00a1-df8d-4611-ac2e-5d6420dbe93b)

It uses [htmx](https://htmx.org/) and [alpine.js](https://alpinejs.dev/start-here) for client side reactivity, on the backend golang echo framework with [pgxpool](https://github.com/jackc/pgx).  
Database used: postgres.

## Setup
```sh
# (using pnpm)
pnpm i

# (sqlc for db codegen)
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

go mod tidy

make dev

# or, in alternative
pnpm watch
go run main.go # (or press F5 in vscode)
```

# Environmental variables
```sh
POSTGRES_DSN=postgres://[username]:[password]@localhost:5432/[db]
```
