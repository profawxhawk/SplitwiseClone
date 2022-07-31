package databases

import (
	"SplitwiseClone/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strings"
)

func InitDB() *gorm.DB {
	connectionString := getDBConnectionString(configs.GetDBConfig())
	log.Println("connectionString", connectionString)
	return Connect(connectionString)
}

func getDBConnectionString(dbConfig *configs.DBConfig) string {
	dataSourceFormat := "{driver}://{username}:{password}@{host}:{port}/{database}"

	r := strings.NewReplacer(
		"{driver}", dbConfig.Driver,
		"{username}", dbConfig.Username,
		"{password}", dbConfig.Password,
		"{host}", dbConfig.Host,
		"{port}", dbConfig.Port,
		"{database}", dbConfig.Database,
	)

	return r.Replace(dataSourceFormat)
}

func Connect(connectionString string) *gorm.DB {
	dbInstance, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Database...")
	return dbInstance
}
