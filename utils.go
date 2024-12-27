package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Db *sql.DB

func LoadEnv() {
	godotenv.Load()
}

func ConnectDb(){
	var err error
	dburi:=os.Getenv("DB_URL")
	fmt.Println(dburi)
	Db,err=sql.Open("mysql",dburi);
	if err!=nil{
		panic(err)
	}
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}