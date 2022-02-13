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

``` sql
CREATE ROLE base_user WITH LOGIN PASSWORD 'users_password';
```

- Create a role for the newly created user, to allow user to create more access

``` sql
ALTER ROLE base_user CREATEDB;
```

- Login as base_user

```
psql postgres -U base_user
```

- Create a new database with this user

``` sql
CREATE DATABASE base_database;
```
