package dbConfig

import (
	"github.com/Croazt/ProjectAkhirPemwebBE/src/utils"
	"github.com/Croazt/ProjectAkhirPemwebBE/src/utils/dbConfig/dbModel"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := utils.GoDotEnvVariable("DATABSE_HOST")
	url := utils.GoDotEnvVariable("DATABSE_URL")
	name := utils.GoDotEnvVariable("DATABASE_NAME")

	var err error
	dsn := host + ":@tcp(" + url + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	if err := DB.AutoMigrate(&dbModel.User{}, &dbModel.Activity{}, &dbModel.Sport{}, &dbModel.SportPlace{}); err != nil {
		panic("Failed to migrate the database")
	}
}
