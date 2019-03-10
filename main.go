package main

import (
	"log"
	"os"

	"github.com/raushan-0822/contacts/config"
	"github.com/raushan-0822/contacts/model"
	"github.com/raushan-0822/contacts/service"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//initialize configmanager
	if err := config.InitConfig(""); err != nil {
		log.Println("Failed initializing configmanager. Error = ", err)
		os.Exit(1)
	}
	config := config.GetConfig()

	if err := model.Init(config.DB["engine"], os.Getenv("MYSQL_URL")); err != nil {
		log.Fatal(err)
	}
	contacts.StartContactAPI(config.Port)
}
