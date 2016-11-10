package config

type productionConfig struct {
}

func (c productionConfig) PostgresURI() string {
	return "postgres://root:CaWUVmZsN29tPD@mitty-db.czsj6nwy3czj.ap-northeast-1.rds.amazonaws.com:5432/mitty?sslmode=disable"
}

func (c productionConfig) PasswordSalt() string {
	return "gBH1SCm3oh6NFPChxGdrCfywEFXVk1sD"
}
