package types

import (
	"fmt"
	"time"
)

const (
	DefaultLocation string = "Europe/Berlin"
)

type RFC3339Time time.Time

func (self RFC3339Time) MarshalJSON() ([]byte, error) {
	locationBerlin, err := time.LoadLocation(DefaultLocation)
	if err != nil {
		return []byte{}, err
	}

	formattedTime := fmt.Sprintf("\"%s\"", time.Time(self).In(locationBerlin).Format(time.RFC3339))

	return []byte(formattedTime), nil
}

func (self *RFC3339Time) UnmarshalJSON(value []byte) error {
	locationBerlin, err := time.LoadLocation(DefaultLocation)
	if err != nil {
		return err
	}

	e := len(value) - 1
	time, err := time.ParseInLocation(time.RFC3339, string(value[1:e]), locationBerlin)
	if err != nil {
		return err
	}

	*self = RFC3339Time(time)

	return nil
}
