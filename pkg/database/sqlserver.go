package database

import (
	"database/sql"
	"fmt"
	"log"
	"main/options"
	"main/pkg/logger"
	"time"

	_ "github.com/alexbrainman/odbc"
)

// DataSource .
const DataSource = "sqlserver://%s:123456@127.0.0.1/database=IPA_%s"

type driver struct {
	db       *sql.DB
	dbDate   string
	username string
}

// DBClient .
type DBClient interface {
	Insert(sql string) error
}

func (d *driver) compareAndResetDatabase() {
	n := time.Now()
	nDbDate := fmt.Sprintf("%v%v", n.Year(), n.Month())
	if nDbDate != d.dbDate {
		d.dbDate = nDbDate
		d.db.Exec(fmt.Sprintf("use IPA_%v", d.dbDate))
	}
}

func (d *driver) Insert(s string) error {
	d.compareAndResetDatabase()
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

	return &driver{
		db:       db,
		username: option.Username,
		dbDate:   "202102",
	}
}
