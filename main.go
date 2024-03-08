package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var (
	fileFlag string
	dirFlag  string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error load .env: %v", err)
	}

	mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: os.Getenv("DBHOST"),
	})
}

func execSqlFile(db *sqlx.DB, path string) {
	_, err := sqlx.LoadFile(db, path)
	if err != nil {
		log.Println("err read file", err)
	}
}

func main() {
	db, err := sqlx.Connect("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("err connect to db: %v", err)
	}

	flag.StringVar(&fileFlag, "file", "default", "execute a sql file")
	flag.StringVar(&dirFlag, "dir", "default", "execute all sql file from a dir")
	flag.Parse()

	if fileFlag != "default" {
		if strings.HasSuffix(fileFlag, ".sql") {
			execSqlFile(db, fileFlag)
		} else {
			log.Fatalf("%s is not valid sql file", fileFlag)
		}
	}
	if dirFlag != "default" {
		entry, err := os.ReadDir(dirFlag)
		if err != nil {
			log.Fatalln(err)
		}
		for _, file := range entry {
			if strings.HasSuffix(file.Name(), ".sql") {
				log.Println("execute:", file.Name())
				path := fmt.Sprintf("%s/%s", dirFlag, file.Name())
				execSqlFile(db, path)
			}
		}
	}
}
