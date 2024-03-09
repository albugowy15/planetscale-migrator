package main

import (
	"crypto/tls"
	"flag"
	"log"
	"os"
	"path/filepath"
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

	err = mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: os.Getenv("DBHOST"),
	})
	if err != nil {
		log.Fatalf("error register tls config: %v", err)
	}
}

func execSqlFile(tx *sqlx.Tx, path string) {
	result, err := sqlx.LoadFile(tx, path)
	if err != nil {
		log.Fatalf("error execute %s: %v", path, err)
	}
	rows, err := (*result).RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%d affected\n", rows)
}

func main() {
	db := sqlx.MustConnect("mysql", os.Getenv("DATABASE_URL"))
	defer db.Close()

	flag.StringVar(&fileFlag, "file", "default", "execute a sql file")
	flag.StringVar(&dirFlag, "dir", "default", "execute all sql file from a dir")
	flag.Parse()

	tx := db.MustBegin()

	if fileFlag != "default" {
		if strings.HasSuffix(fileFlag, ".sql") {
			execSqlFile(tx, fileFlag)
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
				log.Println("execute", file.Name())
				path := filepath.Join(dirFlag, file.Name())
				execSqlFile(tx, path)
			}
		}
	}
	tx.Commit()
}
