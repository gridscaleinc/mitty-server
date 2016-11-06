package config

type developmentConfig struct {
}

func (c developmentConfig) PostgresURI() string {
	return "host=localhost dbname=mitty user=root sslmode=disable"
}
