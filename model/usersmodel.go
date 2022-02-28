package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	usersFieldNames          = builder.RawFieldNames(&Users{}, true)
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "create_time", "update_time"), ",")
	usersRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(usersFieldNames, "id", "create_time", "update_time"))
)

type (
	UsersModel interface {
		Insert(data *Users) (sql.Result, error)
		FindOne(id int64) (*Users, error)
		Update(data *Users) error
		Delete(id int64) error
	}

	defaultUsersModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Users struct {
		Id           int64     `db:"id"`
		RootId       int64     `db:"root_id"` // account  should point to root account if not use phone
		Account      string    `db:"account"`
		Nick         string    `db:"nick"`
		CountryPhone string    `db:"country_phone"`
		Password     int64     `db:"password"`
		Salt         int64     `db:"salt"`
		Introduction string    `db:"introduction"`
		CreateTime   time.Time `db:"create_time"`
		UpdateTime   time.Time `db:"update_time"`
		Avatar       string    `db:"avatar"`
	}
)

func NewUsersModel(conn sqlx.SqlConn) UsersModel {
	return &defaultUsersModel{
		conn:  conn,
		table: `"public"."users"`,
	}
}

func (m *defaultUsersModel) Insert(data *Users) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)", m.table, usersRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Id, data.RootId, data.Account, data.Nick, data.CountryPhone, data.Password, data.Salt, data.Introduction, data.Avatar)
	return ret, err
}

func (m *defaultUsersModel) FindOne(id int64) (*Users, error) {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", usersRows, m.table)
	var resp Users
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Update(data *Users) error {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, usersRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Id, data.RootId, data.Account, data.Nick, data.CountryPhone, data.Password, data.Salt, data.Introduction, data.Avatar)
	return err
}

func (m *defaultUsersModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = $1", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
