package osuapi

import (
	"net/url"
	"strconv"
)

func itos(in int) string {
	return strconv.Itoa(in)
}
func toFuckedUp(a map[string]string) (ret url.Values) {
	ret = url.Values{}
	for k, v := range a {
		ret.Set(k, v)
	}
	return
}
