package config

import "fmt"

// EnvConfigSet ...
type EnvConfigSet interface {
	PostgresURI() string
	PasswordSalt() string
	ESURI() string
}

// CurrentSet ...
var CurrentSet EnvConfigSet

// CurrentEnv ...
var CurrentEnv string

// SetEnvironment ....
func SetEnvironment(envName string) error {
	var configInstance EnvConfigSet

	switch envName {
	case Production:
		configInstance = &productionConfig{}
	case Development:
		configInstance = &developmentConfig{}
	case Docker:
		configInstance = &dockerConfig{}
	}

	if configInstance == nil {
		return fmt.Errorf("No configuration found for %s", envName)
	}

	CurrentSet = configInstance
	CurrentEnv = envName
	return nil
}
