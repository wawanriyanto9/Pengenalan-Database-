package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/YoriDigitalent/Pengenalan-Database1/sql-generic/config"
	"github.com/YoriDigitalent/Pengenalan-Database1/sql-generic/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Println(err)
		return
	}

	db, err := connect(cfg.Database)
	if err != nil {
		log.Println(err)
		return
	}

	//note: hanya digunakan satu per satu
	/*database.InsertCustomer(database.Customer{
		FirstName:    "Sari",
		LastName:     "Kurnia",
		NpwpId:       "npwp987",
		Age:          35,
		CustomerType: "Premiun",
		Street:       "Jalan",
		City:         "Surakarta",
		State:        "Indonesia",
		ZipCode:      "14567",
		PhoneNumber:  "0878341908xxx",
	}, db)*/

	//database.GetCustomers(db)
	//database.DeleteCustomer(2, db)
	database.UpdateCustomer(25, 1, db)

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

func connect(cfg config.Database) (*sql.DB, error) {
	db, err := sql.Open(cfg.Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName, cfg.Config))
	if err != nil {
		return nil, err
	}

	log.Println("db successfully connected")
	return db, nil
}
