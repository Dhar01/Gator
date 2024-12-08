# Aggre-GATOR

**GATOR** (*beware 🐊*) is an RSS feed aggregator for command line, written in Go! Features include:
- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post.

# Install

This project uses Go and PostgresSQL. Please ensure to set them up before proceeding.

1. Install Go
   - Download and install Go from the [official Go website](https://go.dev/dl/).
   - Verify installation:

    ```bash
    go version
    ```

2. Install PostgreSQL
    - Install PostgreSQL from the [official PostgreSQL website](https://www.postgresql.org/download/).
    - Verify Installation:

    ```bash
    psql --version
    ````

3. Install via `go install`

    ```bash
    go install github.com/Dhar01/Gator@latest
    ````

4. On Linux machine, create a configuration file on `$HOME` directory named `.gatorconfig.json`

After installation, run the program using:

```bash
Gator ...
```

# Usage

```bash
$ gator <command> <arguments>
# gator login user
# gator addfeed "Its Foss" https://itsfoss.com/rss/
```

## Available commands

| Command     | Usage                  | Description                        |
| ----------- | ---------------------- | ---------------------------------- |
| `login`     | `login <username>`     | login                              |
| `register`  | `register <username>`  | register                           |
| `reset`     | DANGER                 | DO NOT USE                         |
| `users`     | `users`                | get available users list           |
| `agg`       | `agg <time>`           |                                    |
| `addfeed`   | `addfeed <name> <url>` | add a feed with name and URL       |
| `feeds`     | `feeds`                | get available feeds list           |
| `follow`    | `follow <url>`         | follow a feed using URL            |
| `following` | `following`            | get a list of current feed follows |
| `browse`    |                        |                                    |
| `search`    | N/A                    | N/A                                |
| `help`      | N/A                    | N/A                                |

# TO-DO / Plan

- [ ] Add sorting and filtering options to the `browse` command
- [ ] Add pagination to `browse` command
- [ ] Add concurrency to the `agg` command so that it can fetch more frequently
- [ ] Add a `search` command that allows for fuzzy searching of posts
- [ ] Add bookmarking or liking posts
- [ ] Add a TUI that allows you to select a post in the terminal and view it in a more readable format (*either in the terminal or open in a browser*)
- [ ] Add an HTTP API (*and authentication/authorization*) that allows other users to interact with the service remotely
- [ ] Write a service manager that keeps the `agg` command running in the background and restarts it if it crashes
