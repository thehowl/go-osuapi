package osuapi

// OsuBool is just a bool. It's used for unmarshaling of bools in the API that
// are either `"1"` or `"0"`. thank mr peppy for the memes
//
// You can just use it in `if`s and other memes. Should you need to convert it
// to a native bool, just do `bool(yourOsuBool)`
type OsuBool bool

// UnmarshalJSON converts `"1"` and `1` to true and all other values to false.
func (o *OsuBool) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if dataString == `1` || dataString == `"1"` {
		*o = true
		return nil
	}
	*o = false
	return nil
}

// MarshalJSON does UnmarshalJSON the other way around.
func (o OsuBool) MarshalJSON() ([]byte, error) {
	if o {
		return []byte(`"1"`), nil
	}
	return []byte(`"0"`), nil
}
