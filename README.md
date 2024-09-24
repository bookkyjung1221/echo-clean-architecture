```bash
# clone this repository
git clone https://github.com/bookkyjung1221/echo-clean-architecture.git

#if not have air please install air first
go install github.com/air-verse/air@latest

# start db
docker compose up -d

# run migrate
go run migrate/migrate.go

# start app
air

# stop db
docker compose rm -s -f -v

```
