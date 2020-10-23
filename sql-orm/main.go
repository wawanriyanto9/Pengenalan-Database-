package main

import (
	"fmt"
	"log"

	"github.com/YoriDigitalent/Pengenalan-Database1/sql-orm/config"
	"github.com/YoriDigitalent/Pengenalan-Database1/sql-orm/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Println(err)
		return
	}

	db, err := initDB(cfg.Database)
	if err != nil {
		log.Println(err)
		return
	}

	//digunakan satu per satu

	/*database.InsertCustomer(database.CustomerORM{
		FirstName:    "Kurnia",
		LastName:     "Sari",
		NpwpId:       "456def",
		Age:          18,
		CustomerType: "Premium",
		Street:       "Jalan Sesama",
		City:         "Sukoharjo",
		State:        "Indonesia",
		ZipCode:      "23413",
		PhoneNumber:  "089124679xxx",
	}, db)*/

	//database.GetCustomers(db)
	//database.DeleteCustomer(1, db)
	database.UpdateCustomer(database.CustomerORM{
		PhoneNumber: "081228094xxx",
		City:        "Surakarta",
		Age:         28,
	}, 2, db)

	/*database.InsertAccount(database.AccountORM{
		Balance:     1500,
		AccountType: "Gold",
	}, 1, db)*/
}

func getConfig() (config.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return config.Config{}, err
	}

	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

func initDB(dbConfig config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.Config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&database.CustomerORM{},
		&database.AccountORM{},
	)

	log.Println("db successfully connected")

	return db, nil
}
