## Migration commands

If you're running from Mac OS, from the project root folder please remember to replace `<pwd>` with the actual password, which is in `docker-compose.yml`:

```bash
migrate -source file://db/migrations -database "postgres://postgres:<pwd>@localhost:5433/postgres?sslmode=disable" up
```

The command above should go up all the migration versions, as in `./db/migrations`

```bash
migrate -source file://db/migrations -database "postgres://postgres:<pwd>@localhost:5433/postgres?sslmode=disable" up 1
```

Should go up only one version, from the current one. Please, every time we create a migration, we should at least try to revert it successfully, like this:

```bash
migrate -source file://db/migrations -database "postgres://postgres:<pwd>@localhost:5433/postgres?sslmode=disable" down 1
```

### Creating a new migration script

```bash
migrate create -seq -dir db/migrations -ext sql script_name
```

Creates a new migration script with the proper sequence number, called "script_name.sql"
