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
	s := "select RESULT from IPA_%v.dbo.IPA01 where Model = '%s'"
	rows, err := d.db.Query(fmt.Sprintf(s, nDbDate, barcode))
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var cnt int
	err = rows.Scan(&cnt)
	if err != nil {
		return 0, err
	}
	return cnt, nil
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
