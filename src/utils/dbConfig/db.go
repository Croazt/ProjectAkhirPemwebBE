package dbConfig

import (
	"github.com/Croazt/ProjectAkhirPemwebBE/src/utils"
	"github.com/Croazt/ProjectAkhirPemwebBE/src/utils/dbConfig/dbModel"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := utils.GoDotEnvVariable("DATABASE_HOST")
	// url := utils.GoDotEnvVariable("DATABSE_URL")
	name := utils.GoDotEnvVariable("DATABASE_NAME")
	user := utils.GoDotEnvVariable("DATABASE_USER")
	password := utils.GoDotEnvVariable("DATABASE_PASSWORD")
	port := utils.GoDotEnvVariable("DATABASE_PORT")

	var err error
	dsn := "host=" + host + " user=" + user + "  password=" + password + " dbname=" + name + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	if err := DB.AutoMigrate(&dbModel.User{}, &dbModel.Activity{}, &dbModel.Sport{}, &dbModel.SportPlace{}); err != nil {
		panic("Failed to migrate the database")
	}
}
