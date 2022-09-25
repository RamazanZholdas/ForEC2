package databaseConn

import (
	"fmt"
	"log"
	"os"

	"github.com/RamazanZholdas/APIWithGin/structs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	err error
	Db  *gorm.DB
)

func writeToFile(dns string) {
	file, err := os.Create("dns.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.WriteString(dns)
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectToDB() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("PASSWORD"), os.Getenv("ENDPOINT"), os.Getenv("DBPORT"), os.Getenv("DB"))
	go writeToFile(dns)
	Db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func MigrateModelToDB() {
	Db.AutoMigrate(&structs.Song{})
}
