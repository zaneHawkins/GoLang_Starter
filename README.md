# Go REST API Boilerplate
Golang API boilerplate using GoFiber and PostgreSQL

## Folder Structure

- `/api/v1`
    - `/routes` - All API routes are defined here
    - `/controllers` - For validating requests and calling services
    - `/services` - For business logic, database calls and other services
    - `/middlewares` - For authentication, logging, rate limiting etc.

- `/cmd` - Initializes the fiber app and basic middlewares configuration
- `/config` - For handling configuration/env variables
- `/db` - For handling database connections 
- `/handlers` - For handling responses and db transactions
- `/models` - Auto generated models from database tables using [sqlboiler](https://pkg.go.dev/github.com/volatiletech/sqlboiler/v4@v4.16.1)
- `/types` - For defining custom types that can be used across the app
- `/utils` - For utility functions

- `main.go` - Entrypoint of the app

## Database Guide - PostgreSQL

1. Start Database in Docker `docker-compose up`

2. Shut down database and docker container `docker-compose down`

## Migrate Database Tracking

Golang-migrate is a database version control system written in Golang. This is used to track changes made in sequence to the database using up and down migration files for every change. Up migrations increase the version of the database while down migrations rollback the changes to decrease the version. The `schema_migrations` table stores the current version of the database and if the version is dirty. 

Documentation found here: https://github.com/golang-migrate/migrate.

1. Install migrate CLI - https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
    
    Windows
   1. Install sccop `irm get.scoop.sh | iex`
   2. Install migrate `scoop install migrate`

2. Create new migration files ```migrate create -ext sql -dir ./database/migrations -seq <change_name>```
3. Update the database `migrate -database 'postgresql://user:password@localhost:5432/skillup?sslmode=disable' -path ../database/migrations up`
4. Reset the database  `migrate -database 'postgresql://user:password@localhost:5432/skillup?sslmode=disable' -path ../database/migrations down`

## SQLBoiler Guide

SQLBoiler is a database first ORM. It will automatically generate the database models for the service to use. All files in the `models/` directory are autogenerated.

Documentation found here: https://github.com/volatiletech/sqlboiler.

1. Install sqlboiler `go install github.com/volatiletech/sqlboiler/v4@latest`

2. Install sqlboiler drivers `go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest`

3. Update `sqlboiler.toml` with your database connection information.

4. Run `sqlboiler psql` to generate models based on database tables, this will replace the `models` folder.


## Build Guide

1. Run `go mod tidy` to install all the dependencies.

2. Copy `.env.example` to `.env` and change the values as per your configuration.

3. Run `go build -o ./build/main` to build and run the app.



