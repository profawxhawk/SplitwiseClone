package configs

type DBConfig struct {
	Username   string
	Password   string
	Host       string
	Port       string
	Database   string
	Driver     string
	DBProvider string
}

func NewDBConfig(username string, password string, host string, port string, database string, driver string, dbProvider string) *DBConfig {
	return &DBConfig{Username: username, Password: password, Host: host, Port: port, Database: database, Driver: driver, DBProvider: dbProvider}
}

func GetDBConfig() *DBConfig {
	return NewDBConfig(getEnvWithKey("DB_USER"), getEnvWithKey("DB_PASSWORD"), getEnvWithKey("DB_HOST"), getEnvWithKey("DB_PORT"), getEnvWithKey("DB_NAME"), getEnvWithKey("DB_DRIVER"), getEnvWithKey("DB_PROVIDER"))
}
