# Aggre-GATOR

**GATOR** (*beware 🐊*) is an RSS feed aggregator for command line, written in Go! Features include:
- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post.

# Install

This project uses Go and PostgresSQL. Check them out if you don't have them locally installed. (*to be updated*)

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
