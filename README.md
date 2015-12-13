# go-osuapi

[![docs](https://godoc.org/github.com/thehowl/go-osuapi?status.svg)](https://godoc.org/github.com/thehowl/go-osuapi) [![Build Status](https://drone.io/github.com/thehowl/go-osuapi/status.png)](https://drone.io/github.com/thehowl/go-osuapi/latest)

go-osuapi is an osu! API library for Golang.

## Get started

With this package, mostly everything is very intuitive to do. We recommend to use go-osuapi with a go plugin or a go IDE, so that autocomplete can show up. Once you got autocomplete set up, it is very easy to do anything, as all methods are mostly self-descriptive, or in case they are not they are always "well"-documented (that is because if I don't document them my linter complains about them not being documented all damn day).

Here is an example you can easily copy and paste to hack away:

```go
package main

import (
	"fmt"
	"github.com/thehowl/go-osuapi"
)

func main() {
	c := osuapi.NewClient("Your API key")
	user, err := c.GetUser("peppy", osuapi.Standard)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User %s is rank #%d on osu! standard.", user.Username, user.Rank)
}
```

## Contributing

Contributions are welcome! For conventions and such, we follow a few:

* Always `go fmt` your code.
* We use Test Driven Development to write code. In case you don't like it, that's ok. But please, make sure to have at least one test for the feature you have developed. You can test by putting your osu! API key in the file osukey.txt (it is gitignored so you don't have to worry about it getting accidentally committed) and then running `go test` (optionally with `-v`).
