package account

import (
	"fmt"
)

type TableMetadata struct {
	table   string
	columns []string
}

func (d *TableMetadata) TableName() string {
	return d.table
}

func (d *TableMetadata) Columns() []string {
	return d.columns
}

//a,b,c...
func (d *TableMetadata) ColumnsString() string {
	str := ""
	for i, col := range d.columns {
		if i != 0 {
			str += ", "
		}
		str += col
	}
	return str
}

//a=?,b=?,...
func (d *TableMetadata) ColumnsEqualString() string {
	str := ""
	for i, col := range d.columns {
		if i != 0 {
			str += ", "
		}
		str += (col + " = ?")
	}
	return str
}

func (d *TableMetadata) ColumnsNum() int {
	return len(d.columns)
}

//?,?,?...
func (d *TableMetadata) QuestionMarkString() string {
	str := ""
	for i, _ := range d.columns {
		if i != 0 {
			str += ", "
		}
		str += "?"
	}
	return str
}

type MySQLClientConfig struct {
	Driver   string
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

func (cfg MySQLClientConfig) DataSource() string {
	if cfg.Host == "" {
		cfg.Host = "localhost"
	}
	if cfg.Port == "" {
		cfg.Port = "3306"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}

func (cfg MySQLClientConfig) DriverName() string {
	return cfg.Driver
}
