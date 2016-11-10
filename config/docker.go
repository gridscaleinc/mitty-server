package config

type dockerConfig struct {
}

func (c dockerConfig) PostgresURI() string {
	return "host=postgres dbname=mitty user=root sslmode=disable"
}

func (c dockerConfig) PasswordSalt() string {
	return "FKTaM87v3otln8C3EKaVcL9zhFElaRVC"
}
