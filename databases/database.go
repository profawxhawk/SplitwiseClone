package databases

import (
	"SplitwiseClone/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
)

func InitDB() *gorm.DB {
	connectionString := getDBConnectionString(configs.GetDBConfig())
	return Connect(connectionString)
}

func getDBConnectionString(dbConfig *configs.DBConfig) string {
	dataSourceFormat := "{username}:{password}@tcp({host}:{port})/{database}?parseTime=true&timeout=5s&rejectReadOnly=true"

	r := strings.NewReplacer(
		"{username}", dbConfig.Username,
		"{password}", dbConfig.Password,
		"{host}", dbConfig.Host,
		"{port}", dbConfig.Port,
		"{database}", dbConfig.Database,
	)

	return r.Replace(dataSourceFormat)
}

func Connect(connectionString string) *gorm.DB {
	dbInstance, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
	return dbInstance
}
