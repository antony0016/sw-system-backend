package pgdb

import (
	"fmt"
	"github.com/antony0016/sw-system-backend/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

var DB *gorm.DB

//Init Database initial contain database auto migrate
func Init() {
	if DB != nil {
		return
	}
	Connect()
	core.AutoMigrate(DB)
}

//DsnWithEnv get all data connect query need data by os.Getenv().
func DsnWithEnv() string {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// postgres sql server connect query
	dsnFormat := "host=%s port=%d user=%s password=%s dbname=%s"
	dsn := fmt.Sprintf(dsnFormat, host, port, user, password, dbname)
	return dsn
}

//Connect the func connect to db
func Connect() {
	var (
		err error
	)
	dsn := DsnWithEnv()
	// connect to db and get db instance
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// show connect error
	if err != nil {
		fmt.Println("connect to database fail")
		fmt.Println(err)
	}
}

// "host=%s port=%d user=%s dbname=%s password=%s",
