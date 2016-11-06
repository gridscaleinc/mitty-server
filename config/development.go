package config

type developmentConfig struct {
}

func (c developmentConfig) PostgresURI() string {
	return "host=postgres dbname=mitty user=root sslmode=disable"
}
