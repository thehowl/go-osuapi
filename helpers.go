package osuapi

import (
	"fmt"
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
func checkUsernameType(usernameType string) error {
	if usernameType != "" && usernameType != "id" && usernameType != "string" {
		return fmt.Errorf(`username type is invalid (must be either "string", "id" or empty string)`)
	}
	return nil
}
func checkGamemode(gamemode int) error {
	if gamemode > 3 || gamemode < 0 {
		return fmt.Errorf("passed gamemode is invalid")
	}
	return nil
}
