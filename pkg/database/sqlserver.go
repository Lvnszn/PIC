package database

import (
	"database/sql"
	"fmt"
	_ "github.com/alexbrainman/odbc"
	"log"
	"main/options"
	"main/pkg/logger"
	"time"
)

type driver struct {
	db       *sql.DB
	username string
}

// DBClient .
type DBClient interface {
	Insert(sql string) error
	Select(barcode string) (int, error)
}

func (d *driver) Select(barcode string) (int, error) {
	time.LoadLocation("Asia/Shanghai")
	n := time.Now()
	nDbDate := fmt.Sprintf("%v%.2d", n.Year(), int(n.Month()))
	s := fmt.Sprintf("select Result from IPA_%v.dbo.IPA01 where PShaft = '%s'", nDbDate, barcode)
	logger.Printf("sql is %s", s)
	rows, err := d.db.Query(s)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var cnt int64
		err = rows.Scan(&cnt)
		if err != nil {
			logger.Printf("scan err is %v", err)
			return 0, err
		}

		logger.Printf("scan result is %v", cnt)
		return int(cnt), nil
	}

	return -1, nil
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
	db, err := sql.Open("odbc", "driver={sql server};server=localhost;port=1433;uid=case1;pwd=abc123;database=IPA_202102")
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
