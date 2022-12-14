# shopping-cart-api
Shopping Cart API

## Requirements
- [Golang](https://golang.org/) as main programming language.
- [Go Module](https://go.dev/blog/using-go-modules) for package management.
- [Goose](https://github.com/steinbacher/goose/) as migration tool.
- [Postgresql](https://www.postgresql.org/) as database driver.
- [Docker-compose](https://docs.docker.com/compose/) for running database container locally.
- [Mockery](https://github.com/vektra/mockery/) for generate mockup object

## Setup
### Prepare necessary environment by rename .env.example to .env

### Run database container

```bash
docker-compose up
```

## Run the App
### Get packages

```bash
go get .
```

### Delete unused packages if necessary

```bash
go mod tidy
```

### Update package vendor

```bash
go mod vendor
```

### Build the app

```bash
go build -o bin/shopping-cart -v .
```

### Run the App

```bash
./bin/shopping-cart 
```

## Migration

### Create new migration
```bash
go run app/migration/psql/main.go create AddSomeColumns
```

### Up migration
```bash
go run app/migration/psql/main.go up
```

### Down migration
```bash
go run app/migration/psql/main.go down
```

### Check migration status
```bash
go run app/migration/psql/main.go status
```

## Mockup

### Generate mock

mockery --name=[interface name] --dir=[directory of interface] --output=[directory to write mocks] --filename=[name of generated file] --outpkg=[name of generated package] --structname=[name of generated struct]

```bash
mockery --name=UserRepo --dir=./core/port/ --output=./infrastructure/repo/mock/user --filename=repo.go --outpkg=user --structname=Repo
```

## Unit Test

### Run unit test
```bash
go test -v ./...
```

### Run unit test specific function
```bash
go test -v -run TestProductCategory_Create_Success
```

### Check coverage
```bash
go test -cover ./...
```

### Generate coverage
```bash
go test -coverprofile=coverage ./... 
```