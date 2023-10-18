package database

import (
	"database/sql"
	_ "github.com/microsoft/go-mssqldb"
	"log"
	"main/options"
	"main/pkg/logger"
)

type driver struct {
	db       *sql.DB
	username string
}

// DBClient .
type DBClient interface {
	Insert(sql string) error
}

func (d *driver) Insert(s string) error {
	_, err := d.db.Exec(s)
	return err
}

// NewMssql .
func NewMssql(option *options.Option) DBClient {
	if option.Username == "" {
		option.Username = "case"
	}

	dsn := "server=localhost;port=1433;user id=case1;password=abc123;database=IPA_202309"
	// Create connection pool
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		log.Println("connect to sqlserver fault", err)
		panic(err)
	}
	var (
		sqlversion string
	)
	rows, err := db.Query("select @@version")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err := rows.Scan(&sqlversion)
		if err != nil {
			log.Println(err)
		}
		log.Println(sqlversion)
	}
	logger.Printf("连接数据库成功...")
	return &driver{
		db:       db,
		username: option.Username,
	}
}
