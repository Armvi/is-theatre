package configs

import "errors"

var (
	ErrConfigFileIsEmpty = errors.New("config file is empty")
	// ErrEmptyConfig config struct is empty
	ErrEmptyConfig = errors.New("empty config struct")
	// ErrNoHandlersInConfig unable to unmarshal info about handlers from yaml file
	ErrNoHandlersInConfig = errors.New("unable to get handlers info from config")
)
