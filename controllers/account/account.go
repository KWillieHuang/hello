// Package account provides: 1.user account management, 2.user session management, 3.ldap search function
package account

import (
	"database/sql"
	rlt "hello/models/account"
)

// SysAdminAccount is fixed sysadmin account
var SysAdminAccount = "sysadmin"

// Info account info
type Info struct {
	ID          string
	Domain      string
	UserName    string
	Email       string
	DisplayName string
	IsSysAdmin  bool
}

// UnRegister account and remove all relations,
// may return error if the given account is associated with other business module,
// relationships must be release at first
func UnRegister(accountID string, tx *sql.Tx) error {
	mapper := rlt.NewAccountMapper(tx) //关联至数据库相连结构体
	return mapper.DeleteAccountByID(accountID)
}

// GetList get all account list
func GetList(tx *sql.Tx) (list []Info, err error) {
	mapper := rlt.NewAccountMapper(tx)
	rows, err := mapper.FindAccountAll()
	for _, row := range rows {
		info := Info{}
		info.ID = row.ID
		info.Domain = row.Domain.String
		info.UserName = row.UserName
		info.DisplayName = row.DisplayName
		info.Email = row.Email
		list = append(list, info) //数据写入
	}
	return list, err
}

// GetInfo get account info by id
func GetInfo(accountID string, tx *sql.Tx) (info Info, err error) {
	mapper := rlt.NewAccountMapper(tx)
	row, err := mapper.FindAccountByID(accountID)
	if err != nil {
		return info, err
	}
	info.ID = row.ID
	info.Domain = row.Domain.String
	info.UserName = row.UserName
	info.DisplayName = row.DisplayName
	info.Email = row.Email
	return
}
