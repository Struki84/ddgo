# DuckDuckGo API Query


Small wrapper around DuckDuckGo search API in GoLang. Will return search results based on user input and passed number of wanted results. 


Installation
---
```go
go get https://github.com/Struki84/ddgo
```

Use the query function to perform a search, define the number of returned results. 

```go
results, err := ddgo.Query("2020 FIFA world cup winner?", 10)
```

Results: 

```go
type Result struct {
	Title string
	Info  string
	Ref   string
}
```

Example
---
```go
package main

import (
	"github.com/Struki84/ddgo"
)

func main() {

	results, query := ddgo.Query("2020 FIFA world cup winner?", 5)

	for _, result := range results {
		fmt.Printf("\nTitle: %s \nInfo: %s \nReference: %s \n\n", result.Title, result.Info, result.Ref)
	}
}
```



