package config

type productionConfig struct {
}

func (c productionConfig) PostgresURI() string {
	return "postgres://root:CaWUVmZsN29tPD@mitty.czsj6nwy3czj.ap-northeast-1.rds.amazonaws.com:5432/mitty?sslmode=disable"
}
