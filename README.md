# go-osuapi [![docs](https://godoc.org/gopkg.in/thehowl/go-osuapi.v1?status.svg)](https://godoc.org/github.com/thehowl/go-osuapi) [![Build Status](https://travis-ci.org/thehowl/go-osuapi.svg?branch=master)](https://travis-ci.org/thehowl/go-osuapi)

go-osuapi is an osu! API library for Golang.

## Getting started

Everything is (more or less) well-documented at [godoc](https://godoc.org/gopkg.in/thehowl/go-osuapi.v1) - the methods that interest you most are probably those under [Client](https://godoc.org/gopkg.in/thehowl/go-osuapi.v1#Client). Also, [client-test.go](client-test.go) contains loads of examples on how you can use the package. If you still want to have an example to simply copypaste and then get straight to coding, well, there you go!

```go
package main

import (
	"fmt"
	"gopkg.in/thehowl/go-osuapi.v1"
)

func main() {
	c := osuapi.NewClient("Your API key https://osu.ppy.sh/p/api")
	beatmaps, err := c.GetBeatmaps(osuapi.GetBeatmapsOpts{
		BeatmapSetID: 332532,
	})
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return
	}
	for _, beatmap := range beatmaps {
		fmt.Printf("%s - %s [%s] https://osu.ppy.sh/b/%d\n", beatmap.Artist, beatmap.Title, beatmap.DiffName, beatmap.BeatmapID)
	}
}
```

## Contributing

Contributions are welcome! Here's what you need to know:

* Always `go fmt` your code.
* If you're writing a big and useful feature, make sure to appropriately write tests!
