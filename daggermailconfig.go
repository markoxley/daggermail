package daggermail

import "errors"

type Config struct {
	Host     string
	Port     uint16
	Identity string
	User     string
	Password string
	From     string
}

var configuration *Config

func CreateConfig(host string, port uint16, identity, user, password string) *Config {
	return &Config{
		Host:     host,
		Port:     port,
		Identity: identity,
		User:     user,
		Password: password,
	}
}

func Configure(config *Config) error {
	if config.Host == "" {
		return errors.New("missing host")
	}
	if config.Port == 0 {
		return errors.New("missing port")
	}
	if config.User == "" || config.Password == "" {
		return errors.New("missing credentials")
	}
	if config.From == "" {
		config.From = config.User
	}
	configuration = config
	return nil
}
