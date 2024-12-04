# Aggre-GATOR

**GATOR** (*beware 🐊*) is an RSS feed aggregator for command line, written in Go! Features include:
- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post.

## Available commands

- [x] login (*`login <username>`*)
- [x] register (*`register <username>`*)
- [x] reset (*`reset`*)
- [x] users (*`users`*)
- [x] agg (*`agg`*)
- [x] addfeed (*`addfeed <name> <url>`*)
- [x] feeds (*`feeds`*)
- [x] fetch (*`fetch`*) - will be discarded
- [x] follow (*`follow <url>`*)
- [x] following (*`following`*)