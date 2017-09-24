# go-osuapi [![docs](https://godoc.org/github.com/thehowl/go-osuapi?status.svg)](https://godoc.org/github.com/thehowl/go-osuapi) [![Build Status](https://travis-ci.org/thehowl/go-osuapi.svg?branch=master)](https://travis-ci.org/thehowl/go-osuapi) [![Go Report Card](https://goreportcard.com/badge/github.com/thehowl/go-osuapi)](https://goreportcard.com/report/github.com/thehowl/go-osuapi)

go-osuapi is a Go package to retrieve data from the osu! API.

## Getting started

Everything is (more or less) well-documented at [godoc](https://godoc.org/github.com/thehowl/go-osuapi) - the methods that interest you most are probably those under [Client](https://godoc.org/github.com/thehowl/go-osuapi#Client). Also, [client_test.go](client_test.go) contains loads of examples on how you can use the package. If you still want to have an example to simply copypaste and then get straight to coding, well, there you go!

```go
package main

import (
	"fmt"
	osuapi "github.com/thehowl/go-osuapi"
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

Please note that if you actually want to use this, you should consider vendoring
this using the [dep tool](https://github.com/golang/dep), so that your code
keeps working regardless of any change we might do in this repository.

## I want more than that to explore how it works!

I've made [whosu](https://github.com/thehowl/whosu) for that purpose. Check it out.

## Contributing

Contributions are welcome! Here's what you need to know:

* Always `go fmt` your code.
* If you're writing a big and useful feature, make sure to appropriately write tests!
