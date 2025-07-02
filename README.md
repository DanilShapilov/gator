# Gator
## RSS aggregator (boot.dev project)
Works with:
 - postgres 16.9
 - go 1.24.1

Create `~/.gatorconfig.json` with following content, don't forget to change postgres credentials and port if needed:
```json
{
    "db_url":"postgres://postgres:assword@localhost:5432/gator?sslmode=disable",
    "current_user_name":"kahya"
}
```

 When in project directory you can install it with `go install` and then use like so:  
`gator register <user_name>` - creates user with given name and sets him as current user  
`gator login <user_name>` - sets current user  
`gator reset` - resets database  
`gator users` - lists users  
`gator agg <time_between_reqs>` - scrapes feeds, adding posts to local DB    
`gator browse [limit=2]` - list posts from local DB  
`gator addfeed <name> <url>` - add feed  
`gator feeds` - list feeds  
`gator follow <url>` - follow feed  
`gator unfollow <url>`- unfollow feed  
`gator following` - list followed feeds  