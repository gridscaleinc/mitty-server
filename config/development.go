package config

type developmentConfig struct {
}

func (c developmentConfig) PostgresURI() string {
	return "postgres://root:mpNffadJrfnWpvZxnrZz2Zjz@52.196.151.53:5432/mitty_db?sslmode=disable"
}

func (c developmentConfig) PasswordSalt() string {
	return "tJhFkhVNR3vJclPb56V5aA9n37TxXT4O"
}
