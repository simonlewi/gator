# Gator - RSS Feed Aggregator CLI

A command-line RSS feed aggregator built in Go that helps you follow and aggregate blog posts from your favorite RSS feeds.

## Prerequisites

- Go 1.16 or higher
- PostgreSQL 12 or higher
- Basic knowledge of command line interfaces

## Installation

1. Install the CLI:
```bash
go install github.com/simonlewi/gator@latest
```

2. Create a config file at `~/.gatorconfig.json`:
```json
{
    "db_url": "postgresql://username:password@localhost:5432/gator?sslmode=disable",
    "current_username": ""
}
```

Replace `username`, `password`, and database name as needed.

## Usage

### Basic Commands

- `gator register <username>` - Create a new user account
- `gator login <username>` - Login as a user
- `gator addfeed <name> <url>` - Add a new RSS feed
- `gator feeds` - List all available feeds
- `gator following` - List feeds you're following
- `gator follow <url>` - Follow a feed
- `gator unfollow <url>` - Unfollow a feed
- `gator browse` - View posts from feeds you follow

### Aggregation Mode

To continuously fetch new posts from followed feeds:

```bash
gator agg 1m
```

This will check for new posts every minute. You can adjust the duration (e.g., `10s`, `5m`, `1h`).

### Example Workflow

1. Create an account:
```bash
gator register johndoe
```

2. Add and follow a feed:
```bash
gator addfeed "Go Blog" "https://go.dev/blog/feed.atom"
```

3. View your followed feeds:
```bash
gator following
```

4. Start aggregating posts:
```bash
gator agg 5m
```

5. Browse posts:
```bash
gator browse
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

MIT