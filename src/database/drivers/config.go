package drivers

import "errors"

var ErrConfigNotSet = errors.New("datastore config is not set")

type Config struct {
	DBName string
	DSName string
	URL    string
}
