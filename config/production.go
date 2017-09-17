package config

type productionConfig struct {
}

func (c productionConfig) PostgresURI() string {
	return "postgres://mitty_user:PxeFKA9nXawKhyrFCi2Ajrenyzkocy@13.115.50.60:5432/mitty_db?sslmode=disable"
}

func (c productionConfig) PasswordSalt() string {
	return "gBH1SCm3oh6NFPChxGdrCfywEFXVk1sD"
}

func (c productionConfig) ESURI() string {
	return "http://52.197.56.194:9200"
}
