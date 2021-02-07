package database

import (
	"database/sql"
	"log"
	"main/options"
	"main/pkg/logger"

	_ "github.com/alexbrainman/odbc"
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

	// Create connection pool
	db, err := sql.Open("odbc", "driver={sql server};server=localhost;port=1433;uid=case;pwd=123456;database=IPA_202102")
	if err != nil {
		logger.Printf("err is %v", err)
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
