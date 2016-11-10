package config

type developmentConfig struct {
}

func (c developmentConfig) PostgresURI() string {
	return "host=localhost dbname=mitty user=root sslmode=disable"
}

func (c developmentConfig) PasswordSalt() string {
	return "tJhFkhVNR3vJclPb56V5aA9n37TxXT4O"
}
