package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	asiaDB    *gorm.DB
	americaDB *gorm.DB
)
var err error

func InitDB() {

	asiaDBUser := os.Getenv("DB_ASIA_USER")
	asiaDBPassword := os.Getenv("DB_ASIA_PASSWORD")
	asiaDBHost := os.Getenv("DB_ASIA_HOST")
	asiaDBPort := os.Getenv("DB_ASIA_PORT")
	asiaDBName := os.Getenv("DB_ASIA_NAME")

	americaDBUser := os.Getenv("DB_AMERICA_USER")
	americaDBPassword := os.Getenv("DB_AMERICA_PASSWORD")
	americaDBHost := os.Getenv("DB_AMERICA_HOST")
	americaDBPort := os.Getenv("DB_AMERICA_PORT")
	americaDBName := os.Getenv("DB_AMERICA_NAME")

	asiaDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", asiaDBHost, asiaDBUser, asiaDBPassword, asiaDBName, asiaDBPort)
	americaDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", americaDBHost, americaDBUser, americaDBPassword, americaDBName, americaDBPort)

	asiaDB, err = gorm.Open(postgres.Open(asiaDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the Asia database: %v", err)
	}

	americaDB, err = gorm.Open(postgres.Open(americaDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the America database: %v", err)
	}

	// Ensure the connections are established
	sqlAsiaDB, err := asiaDB.DB()
	if err != nil {
		log.Fatalf("Error getting the Asia database object: %v", err)
	}
	err = sqlAsiaDB.Ping()
	if err != nil {
		log.Fatalf("Error pinging the Asia database: %v", err)
	}

	sqlAmericaDB, err := americaDB.DB()
	if err != nil {
		log.Fatalf("Error getting the America database object: %v", err)
	}
	err = sqlAmericaDB.Ping()
	if err != nil {
		log.Fatalf("Error pinging the America database: %v", err)
	}

	fmt.Println("Asia and America Database connections successful")
}

func GetDB(region string) *gorm.DB {
	switch region {
	case "Asia":
		return asiaDB
	case "America":
		return americaDB
	default:
		log.Fatalf("Invalid region: %s", region)
		return nil
	}
}
