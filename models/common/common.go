package common

import (
	"database/sql"
	rlt "hello/models/account"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var RoutineChan chan int
var Log *logs.BeeLogger

// SQLTx struct
type SQLTx struct {
	Tx          *sql.Tx
	DB          *sql.DB
	RoutineChan chan int
}

// Close close sql db and commit tx
func (tx *SQLTx) Close() error {
	defer func() {
		<-tx.RoutineChan
	}()
	err := tx.DB.Close()
	if err != nil {
		return err
	}
	err = tx.Tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func GetTx() (*SQLTx, error) {
	rltcfg := rlt.MySQLClientConfig{
		Driver:   beego.AppConfig.String("mysqldriver"),
		Host:     beego.AppConfig.String("mysqlhost"),
		Port:     beego.AppConfig.String("mysqlport"),
		Database: beego.AppConfig.String("mysqldatabase"),
		User:     beego.AppConfig.String("mysqluser"),
		Password: beego.AppConfig.String("mysqlpassword"),
	}
	db, err := sql.Open(rltcfg.DriverName(), rltcfg.DataSource())
	if err != nil {
		Log.Error(err.Error())
		return nil, err
	}
	db.SetMaxOpenConns(2000)
	tx, err := db.Begin()
	if err != nil {
		Log.Error(err.Error())
		return nil, err
	}
	if RoutineChan == nil {
		RoutineChan = make(chan int, 2000)
	}
	RoutineChan <- 0
	return &SQLTx{tx, db, RoutineChan}, nil
}
