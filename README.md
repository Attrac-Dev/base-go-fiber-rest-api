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

# added the export path into the watch method
watch:
    export PATH=$$PATH:$$GOPATH/bin
    reflex -s -r '\.go$$' make run'
```

In order to get reflex working correctly, we need to do `export PATH=$PATH:$GOPATH/bin` to update the path for the files.

 `make watch` will the be the command used to run the 'development' server with hot reloading.

---

## Youtube video for how to setup database locally

[video link](https://www.youtube.com/watch?v=fGOsgMcTP2I)

- PostgreSQL database was installed locally with Homebrew

- start PostgreSQL with `brew services postgresql`

- stop PostgreSQL with `brew services stop postgresql'

- to start the PostgreSQL CLI use `psql postgres`

- to stop the CLI interface type 'exit' or '\q' into the command line

- will need to set up some alias in my system to make it easier to remember these

- Create a new user for the database

---
[gist with basic Postgres commands for creating user, etc.](https://gist.github.com/phortuin/2fe698b6c741fd84357cec84219c6667)

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

`.env` file string can look like:

``` bash
PG_CONNECTION_STRING=postgres://myuser@localhost/mydatabase
```

### - Don't forget to add this to the .gitignore

### `.env`

*** updated the DB_PORT to jive with the Docker postgres:alpine info later on

``` bash
DB_HOST= localhost
DB_NAME= base_database
DB_USER= base_user
DB_PASSWORD= users_password
DP_PORT= 5555
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

### Add `gorm` and `postgres driver` by running

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

---

There was a naming conflict with the original module name

Had to revert the module name in `go.mod` from the 'Attrac-Dev' github, to the 'skyler-saville' github, before adding the database connection file to the `main.go`

---

## Adding Data Models

The original tutorial was using `notes` as the only model in the project. In order to keep this repo pretty generic, I am going to be making this as simplified as possible, when it comes to the data models. In the future, I may fork this repo to use as the starting point for more complex APIs.

### Make a folder called `internal`

> All of the internal logic for the API will reside in this space. Within the `internal` directory, back a sub-directory with the name `models`, which will  contain all of the different data models. Within the `models` directory add a `.go`  file with the name of the model (in this repo it is just called `model`).

## Auto Migrations

> GORM supports auto migrations, whenever a change is made to a models' struct (add  a column, change a type, add and index, etc.) and restart the server, the changes will be reflected in the database automatically.

## ADDED DOCKER

Followed video on [Youtube](https://www.youtube.com/watch?v=aHbE3pTyG-Q)

``` bash
docker run --name postgres-test -e POSTGRES_PASSWORD=password -d -p 5555:5432 postgres:alpine
```

on local machine, connect to the container with the locally installed version of psql:

``` bash
psql -h localhost -p 5555 -U postgres
```

on local machine, bash into the container:

``` bash
docker exec -it postgres-test bash
```

then login as the super postgres user:

``` bash
psql -U postgres
```

create a new user:

``` bash
CREATE ROLE base_user WITH LOGIN PASSWORD 'users_password';
```

set users role:

``` bash
ALTER ROLE base_user CREATEDB;
```

create the database and set the new user as the owner:

``` bash
CREATE DATABASE base_database WITH OWNER base_user
```

login with user and connect to the new database:

``` bash
\c base_database base_user
```

## Changed from Reflex to Nodemon

You have to install nodemon globally on your machine using:

``` bash
npm i -g nodemon
```

Then have to update the makefile to utilize nodemon instead of reflex

``` makefile
watch: 
 nodemon --exec "go run" main.go
```

Now you can start the server in "dev" mode by running

``` bash
make watch
```

Then you can test the endpoints of the router with Postman, Insomnia, etc.

## Task Model

- ID uuid.UUID

- Title     *string

- Subtitle  string

- Text      string

- CompletedOnDate   time.Time

## API endpoints for "tasks"

- **POST**   0.0.0.0:3000/api/tasks  _(create task)_

- **GET**   0.0.0.0:3000/api/tasks  _(get all tasks)_

- **GET**   0.0.0.0:3000/api/tasks/:taskID  _(get one task)_

- **PUT**   0.0.0.0:3000/api/tasks/:taskID  _(update a task)_

- **DELETE**   0.0.0.0:3000/api/tasks/:taskID  _(delete a task)_
