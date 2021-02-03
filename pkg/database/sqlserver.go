package database

import (
	"database/sql"
	"fmt"
	"main/options"
	"main/pkg/logger"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

// DataSource .
const DataSource = "sqlserver://%s@192.168.0.10/IPA_%s?sslmode=disable"

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

		db, err := sql.Open("mssql", fmt.Sprintf(DataSource, d.username, d.dbDate))
		if err != nil {
			logger.Printf("init database fail %v", err)
			return
		}
		d.db.Close()
		d.db = db
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
		option.Username = "ADMIN"
	}
	db, err := sql.Open("mssql", fmt.Sprintf(DataSource, option.Username, "202102"))
	if err != nil {
		logger.Printf("err is %v", err)
		panic(err)
	}
	logger.Printf("connect to %v success", fmt.Sprintf(DataSource, option.Username, "202102"))
	return &driver{
		db:       db,
		username: option.Username,
		dbDate:   "202102",
	}
}
