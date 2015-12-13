package osuapi

import (
	"time"
)

// MySQLDate is a wrapper for time.Time that can get a date from a JSON response.
type MySQLDate time.Time

// UnmarshalJSON takes some JSON data and does some magic to transform it into a native time.Time.
func (m *MySQLDate) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if dataString == "null" {
		m = nil
	}
	inTimeLib, err := time.Parse(`"2006-01-02 15:04:05"`, dataString)
	if err != nil {
		return err
	}
	*m = MySQLDate(inTimeLib)
	return nil
}

// GetTime transforms a MySQLDate into a native time.Time.
func (m MySQLDate) GetTime() time.Time {
	return time.Time(m)
}

// String converts a MySQLDate to a string.
func (m MySQLDate) String() string {
	return m.GetTime().Format("2006-01-02 15:04:05")
}
