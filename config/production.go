package config

type productionConfig struct {
}

func (c productionConfig) PostgresURI() string {
	return "host=postgres dbname=mitty user=root sslmode=disable"
}
