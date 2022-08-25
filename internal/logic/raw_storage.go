package logic

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// MYDB is a DB wrapper which logs the executed sql queries and their
// duration.
type MYDB struct {
	*sql.DB
	DbType string
}

// Query logs the queries executed by the Query method.
func (db *MYDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	start := time.Now()
	rows, err := db.DB.Query(query, args...)
	logQuery(query, time.Since(start), args...)
	return rows, err
}

// Exec logs the queries executed by the Exec method.
func (db *MYDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	res, err := db.DB.Exec(query, args...)
	logQuery(query, time.Since(start), args...)
	return res, err
}

func filterbyte(fb string) (rfb string) {
	rfb = fb
	for _, symbol := range []string{"-", ";", ",", "\\", "/", "[", "]", "{", "}", "%", "@", "*", "!", "'"} {
		rfb = strings.Replace(rfb, symbol, "_", -1)
	}
	rfb = `TB_` + rfb
	return
}

// CreateDevTable creat new device table
func (db *MYDB) CreateDevTable(tbname string) (err error) {
	err = errors.New("database not init")
	var column string
	if db == nil {
		return
	}
	tbname = filterbyte(tbname)
	start := time.Now()
	if db.DbType == "mysql" {
		column = ` (timestamp TIMESTAMP WITH TIME ZONE PRIMARY KEY,cmdtype TEXT, jdoc jsonb);`
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS ` + tbname + column)
		log.WithFields(log.Fields{
			"create":   tbname,
			"duration": time.Since(start),
		}).Debug("sql create table executed")
	}
	if db.DbType == "sqlite3" {
		column = ` (timestamp integer not null primary key, cmdtype text, jdoc text);`
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS ` + tbname + column)
		log.WithFields(log.Fields{
			"create":   tbname,
			"duration": time.Since(start),
		}).Debug("sql create table executed")
	}
	return
}

// TxLogger logs the executed sql queries and their duration.
type TxLogger struct {
	*sql.Tx
}

// Query logs the queries executed by the Query method.
func (q *TxLogger) Query(query string, args ...interface{}) (*sql.Rows, error) {
	start := time.Now()
	rows, err := q.Tx.Query(query, args...)
	logQuery(query, time.Since(start), args...)
	return rows, err
}

// Exec logs the queries executed by the Exec method.
func (q *TxLogger) Exec(query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	res, err := q.Tx.Exec(query, args...)
	logQuery(query, time.Since(start), args...)
	return res, err
}

func logQuery(query string, duration time.Duration, args ...interface{}) {
	log.WithFields(log.Fields{
		"query":    query,
		"args":     args,
		"duration": duration,
	}).Debug("sql query executed")
}

// InsertMap ....
func InsertMap(db *sql.DB, tbname string, m map[string]string) (err error) {
	err = errors.New("database not init")
	if db == nil {
		return
	}
	column := new(bytes.Buffer)
	value := new(bytes.Buffer)
	column.WriteString("(")
	value.WriteString("VALUES (")
	nl := len(m)
	i := 1
	format := "%s,"
	formatv := "'%s',"
	for k, v := range m {
		if i == nl {
			format = "%s"
			formatv = "'%s'"
		}
		i++
		column.WriteString(fmt.Sprintf(format, k))
		value.WriteString(fmt.Sprintf(formatv, v))
	}
	column.WriteString(")")
	value.WriteString(")")
	strinsert := `INSERT INTO ` +
		tbname +
		" " +
		column.String() +
		" " +
		value.String() + ";"
	fmt.Println(strinsert)
	stmt, err := db.Prepare(strinsert)
	defer stmt.Close()
	if err != nil {
		return errors.Wrap(err, "get stmt error")
	}
	_, err = stmt.Exec()
	return
}

func queryWithCondition(db *sql.DB, table, id, idname, startt, endt string, offset, limit int) ([]any, error) {
	//thisoffset := offset * limit
	queryString := fmt.Sprintf("select * from %s where %s='%s' and %s between ? and ?;", table, id, idname, "timestamp")
	var err error
	rows, _ := db.Query(queryString, startt, endt)
	col, _ := rows.Columns()
	vs := make([]any, len(col))
	rawb := make([]sql.RawBytes, len(col))
	for i := range rawb {
		vs[i] = &rawb[i]
	}
	for rows.Next() {
		rows.Scan(vs...)
		log.Infof("%s", rawb)
	}
	rows.Close()
	return nil, err
}
