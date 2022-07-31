package configs

type DBConfig struct {
	Username   string
	Password   string
	Host       string
	Port       string
	Database   string
	DBProvider string
}

func NewDBConfig(username string, password string, host string, port string, database string, dbProvider string) *DBConfig {
	return &DBConfig{Username: username, Password: password, Host: host, Port: port, Database: database, DBProvider: dbProvider}
}

func GetDBConfig() *DBConfig {
	return NewDBConfig(getEnvWithKey("DB_USER"), getEnvWithKey("DB_PASSWORD"), getEnvWithKey("DB_HOST"), getEnvWithKey("DB_PORT"), getEnvWithKey("DB_NAME"), getEnvWithKey("DB_PROVIDER"))
}
