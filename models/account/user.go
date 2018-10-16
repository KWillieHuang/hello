package account

import (
	"database/sql"
	"fmt"
)

/***************************************************************************
Account Mapper

SQL:
	--
	-- Definition for table account
	--
	CREATE TABLE account (
	  id VARCHAR(32) NOT NULL,
	  domain VARCHAR(32)
	  user_name VARCHAR(128) NOT NULL,
	  display_name VARCHAR(128) NOT NULL,
	  email VARCHAR(128) NOT NULL,
	  PRIMARY KEY (id),
	  UNIQUE INDEX user_name (user_name)
	)
	ENGINE = INNODB
	AVG_ROW_LENGTH = 16384
	CHARACTER SET utf8
	COLLATE utf8_general_ci;
***************************************************************************/

type AccountRow struct {
	ID          string
	Domain      sql.NullString
	UserName    string
	DisplayName string
	Email       string
}

type AccountMapper struct {
	Tx       *sql.Tx
	Metadata TableMetadata
}

func (m *AccountMapper) fields(row *AccountRow) []interface{} {
	return []interface{}{
		&row.ID,
		&row.Domain,
		&row.UserName,
		&row.DisplayName,
		&row.Email,
	}
}

func (m *AccountMapper) CreateAccount(id string, domain sql.NullString, userName string, displayName string, email string) (*AccountRow, error) {
	sqlInsert := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s);`,
		m.Metadata.TableName(),
		m.Metadata.ColumnsString(),
		m.Metadata.QuestionMarkString())
	res, err := m.Tx.Exec(sqlInsert, id, domain, userName, displayName, email)
	if err != nil {
		return nil, err
	}

	if num, err := res.RowsAffected(); err != nil || num != 1 {
		return nil, fmt.Errorf("RowsAffected: %d, %s", num, err.Error())
	}

	row := &AccountRow{}
	sqlSelect := fmt.Sprintf(`SELECT %s FROM %s WHERE id = ?`, m.Metadata.ColumnsString(), m.Metadata.TableName())
	err = m.Tx.QueryRow(sqlSelect, id).Scan(m.fields(row)...)
	if err != nil {
		return nil, err
	}
	return row, nil
}

//FindAccount
func (m *AccountMapper) FindAccountByID(id string) (*AccountRow, error) {
	row := &AccountRow{}
	sqlSelect := fmt.Sprintf(`SELECT %s FROM %s WHERE %s = ?`, m.Metadata.ColumnsString(), m.Metadata.TableName(), id)
	err := m.Tx.QueryRow(sqlSelect, id).Scan(m.fields(row)...)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (m *AccountMapper) DeleteAccountByID(id string) error {
	sqlSelect := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, m.Metadata.TableName())
	res, err := m.Tx.Exec(sqlSelect, id)
	if err != nil {
		return err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if num != 1 {
		return fmt.Errorf("AffectedRows:%d, now rows found", num)
	}
	return nil
}

func (m *AccountMapper) FindAccountAll() ([]AccountRow, error) {
	sql := fmt.Sprintf(`SELECT %s FROM %s`, m.Metadata.ColumnsString(), m.Metadata.TableName())
	rows, err := m.Tx.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	slice := make([]AccountRow, 0)
	for rows.Next() {
		row := AccountRow{}
		if err := rows.Scan(m.fields(&row)...); err != nil {
			return nil, err
		}
		slice = append(slice, row)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return slice, nil
}

func NewAccountMapper(tx *sql.Tx) *AccountMapper {
	mapper := &AccountMapper{
		Tx: tx,
		Metadata: TableMetadata{
			table: "account",
			columns: []string{
				"id",
				"domain",
				"user_name",
				"display_name",
				"email",
			},
		},
	}
	return mapper
}
