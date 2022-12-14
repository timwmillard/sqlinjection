# SQL Injection

This is a sample project that has a deliberate SQL exploit built in.  See if you can find it.

## Problem

Run a command that will delete all data from the `person` table.

---

## Setup

Setup the database and build the program.

### Database Environment

Ensure a Postgres database is running.

Create a `.env` file in the root directory and set the database variables.

Example:
```
DB_USER=user
DB_PASSWORD=password
DB_NAME=my_database
DB_HOST=localhost
DB_PORT=5434
```

### Migrate the Database

Ensure you have `psql` installed.  Then run:

```sh
make migrate
```

### Build

```sh
make
```

---

## Run the Program

This will show the available commands.

```sh
./person
```


### Insert some data

```sh
./person create Tim Millard
./person create John Smith
```

### View the data

```sh
./person list
```

### Update the data

```sh
./person update 1 Timothy Millard
```

---

## Hack

Now hack away and find the exploit.
