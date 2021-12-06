package blogmgr

import (
	bloglog "blog/log"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var logger bloglog.Logger
var db *sql.DB
var bInit = false

func InitAPI(filelogger *bloglog.Logger) error {
	if bInit {
		return nil
	}
	logger = *filelogger
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog")
	if err != nil {
		logger.ErrErr("sql.open failed", err)
		return err
	}
	db.SetConnMaxLifetime(100)
	db.SetMaxIdleConns(100)
	bInit = true
	return nil
}
