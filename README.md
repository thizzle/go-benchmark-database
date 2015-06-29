# Go Database Library Benchmarks

A suite of benchmarks for database libraries in Go. Currently benchmarks:

* `database/sql` package (built-in)
* [`database/sqlx`](https://github.com/jmoiron/sqlx)

## Running via Docker

The easiest way to run these benchmarks is via [Docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/) (previously known as fig). You'll want both installed and available on your `$PATH`.

1. `$ docker-compose build` to build the required Docker image (just for the application).

2. `$ docker-compose up -d pgm` to start a database server instance (postgresql).

3. `$ docker-compose run app goose up` (only once) to run the database migrations, and bring the database up-to-date.

4. `$ docker-compose run app` to run the tests/benchmarks.

