# Go Fiber Restful API

## Testing out the Go Fiber framework, which is very similar to Express.js

every import needs to be added using the:

``` go
go get -u <some github url>
```

This follows along with the tutorial on [dev.to](https://dev.to/percoguru/getting-started-with-apis-in-golang-feat-fiber-and-gorm-2n34).

The tutorial used "notes" as the main api resource, but this is going to be a repo made with the purpose of cloning it to shorten API development times. The resource here is going to be a little more generic.

## Items on the agenda

- add a Makefile
- add a database (PostgreSQL)
- add the resource Models
- and some environment variables (.env file)

To enable hot reloading, we will need to add a `Makefile` to the root directory.

- Install reflex
`go install github.com/cespare/reflex@latest`

- Add commands to the Makefile:

``` Makefile
build:
    go build -o server main.go

run: build
    ./server

watch:
    reflex -s -r '\.go$$' make run'
```

In order to get reflex working correctly, we need to do `export PATH=$PATH:$GOPATH/bin` to update the path for the files.

 `make watch` will the be the command used to run the 'development' server with hot reloading.

<hr>

## Youtube video for how to setup database locally

[video link](https://www.youtube.com/watch?v=fGOsgMcTP2I)

- PostgreSQL database was installed locally with Homebrew

- start PostgreSQL with `brew services postgresql`

- stop PostgreSQL with `brew services stop postgresql'

- to start the PostgreSQL CLI use `psql postgres`

- to stop the CLI interface type 'exit' or '\q' into the command line

- will need to set up some alias in my system to make it easier to remember these

- Create a new user for the database

``` pgsql
CREATE ROLE base_user WITH LOGIN PASSWORD 'users_password';
```

- Create a role for the newly created user, to allow user to create more access

``` pgsql
ALTER ROLE base_user CREATEDB;
```

- Login as base_user

``` pgsql
psql postgres -U base_user
```

- Create a new database with this user

``` pgsql
CREATE DATABASE base_database;
```

## Set up the .env with the database creds

### - Don't forget to add this to the .gitignore

### `.env`

```
DB_HOST= localhost
DB_NAME= base_database
DB_USER= base_user
DB_PASSWORD= users_password
DP_PORT= 5432
```

### Now use `go get` to add the `godotenv` module

``` go
go get github.com/joho/godotenv
```

### Create `config.go`

``` go
package config

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
    // load .env file
    err := godotenv.Load(".env")
    if err != nil {
        fmt.Print("Error loading .env file")
    }
        // Return the value of the variable
    return os.Getenv(key)
}
```

## Connecting to the PostgreSQL database

### Add `grom` and `postgres driver` by running

``` go
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

## `/database/connect.go`

``` go
package database

import (
    "fmt"
    "log"
    "strconv"

    "github.com/percoguru/notes-api-fiber/config"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
    var err error
    p := config.Config("DB_PORT")
    port, err := strconv.ParseUint(p, 10, 32)

    if err != nil {
        log.Println("Idiot")
    }

    // Connection URL to connect to Postgres Database
    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
    // Connect to the DB and initialize the DB variable
    DB, err = gorm.Open(postgres.Open(dsn))

    if err != nil {
        panic("failed to connect database")
    }

    fmt.Println("Connection Opened to Database")
}
```

### `connect.go` imports the package config. It looks for the package inside the folder `./config`
