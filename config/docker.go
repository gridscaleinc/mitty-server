package config

type dockerConfig struct {
}

func (c dockerConfig) PostgresURI() string {
	return "postgres://root:mpNffadJrfnWpvZxnrZz2Zjz@dev.mitty.co:5432/mitty_db?sslmode=disable"
}

func (c dockerConfig) PasswordSalt() string {
	return "FKTaM87v3otln8C3EKaVcL9zhFElaRVC"
}
