# gator

## Requirements

- [Go](https://golang.org/doc/install) (version 1.18+ recommended)
- [PostgreSQL](https://www.postgresql.org/download/)

## Installation

Install the gator CLI using Go:

```
go install github.com/tbone317/gator@latest
```
This will place the `gator` binary in your `$GOPATH/bin` or `$HOME/go/bin` directory.

## Configuration

Create a `.gatorconfig.json` file in your home directory or project root. Example:

```
{
	"db_url": "postgres://username:password@localhost:5432/gatordb?sslmode=disable",
	"current_user_name": "yourusername"
}
```

Update `db_url` with your Postgres connection string and set your username.

## Running the Program

Run the CLI from your terminal:

```
gator <command> [args]
```

## Example Commands

- `gator login <username>`: Log in as a user
- `gator browse`: Browse available feeds
- `gator follow <feed_url>`: Follow a new feed
- `gator unfollow <feed_url>`: Unfollow a feed
- `gator reset`: Reset your user state

For more commands, see the source or run `gator help`.