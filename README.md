## Gira Client

This is the Go client, for the Gira API.

### How to use
#### Download via Go Modules
```
$ go get github.com/gira-games/client
```
#### Use it in your application like that
```go
import "github.com/gira-games/client"

cl := client.New("https://api.gira.com")
resp, err := cl.GetGames(context.Background(), &client.GetGamesRequest{
    Token: "my-token"
})
// use resp and err
```

### License
This work is licensed under MIT license. For more info see [LICENSE.md](LICENSE.md)
