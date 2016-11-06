package config

type dockerConfig struct {
}

func (c dockerConfig) PostgresURI() string {
	return "host=postgres dbname=mitty user=root sslmode=disable"
}
