package logic

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
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

// InsertDevJdoc insert device value
func (db *MYDB) InsertDevJdoc(tbname, cmdtype string, value interface{}) (err error) {
	err = errors.New("database not init")
	if db == nil {
		return
	}
	tbname = filterbyte(tbname)
	items := `(timestamp,cmdtype,jdoc) VALUES($1,$2,$3) RETURNING timestamp`
	if db.DbType == "sqlite3" {
		items = `(timestamp,cmdtype,jdoc) VALUES($1,$2,$3)`
	}
	strinsert := `INSERT INTO ` +
		tbname +
		items
	stmt, err := db.Prepare(strinsert)
	defer stmt.Close()
	if err != nil {
		return errors.Wrap(err, "get stmt error")
	}
	//mymsg, err := simplejson.NewJson(xxjson)
	//logsb, _ := mymsg.EncodePretty()
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	jdoc, err := json.Marshal(value)
	if err != nil {
		return errors.Wrap(err, "marshal json error")
	}
	if db.DbType == "mysql" {
		_, err = stmt.Exec(time.Now(), cmdtype, string(jdoc))
	}
	if db.DbType == "sqlite3" {
		_, err = stmt.Exec(time.Now().Unix(), cmdtype, string(jdoc))
	}
	return
}

func (db *MYDB) QueryDevJdoc(tbname, cmdtype, since, until string) (ret []map[string]interface{}, err error) {
	if db == nil {
		return nil, errors.New("mysql is not connect or connect error")
	}
	tbname = filterbyte(tbname)
	defer func() {
		if Err := recover(); Err != nil {
			log.Errorf("panic  error : (%s)", Err)
			err = fmt.Errorf("panic  error : %s", Err)
		}
	}()
	var condition, strquery string
	condition = ` WHERE timestamp > '` +
		since +
		` +08:00:00' and timestamp < '` +
		until +
		` +08:00:00'` +
		`and cmdtype LIKE '` +
		`%` +
		cmdtype +
		`%` +
		`'`
	strquery = `SELECT timestamp AT TIME ZONE 'Asia/Shanghai',cmdtype,jdoc FROM ` +
		tbname + condition + `;`
	rows, qerr := db.Query(strquery)
	if qerr == nil {
		for rows.Next() {
			var rjs, cmt string
			if db.DbType == "mysql" {
				var timesp time.Time
				if serr := rows.Scan(&timesp, &cmt, &rjs); serr == nil {
					sjrjs, _ := simplejson.NewJson([]byte(rjs))
					ret = append(ret, map[string]interface{}{
						"timestamp": timesp,
						"cmdtype":   cmt,
						"value":     sjrjs,
					})
				}
			}
			if db.DbType == "sqlite3" {
				var rjs, cmt string
				var timesp int64
				if serr := rows.Scan(&timesp, &cmt, &rjs); serr == nil {
					sjrjs, _ := simplejson.NewJson([]byte(rjs))
					ret = append(ret, map[string]interface{}{
						"timestamp": time.Unix(timesp, 0),
						"cmdtype":   cmt,
						"value":     sjrjs,
					})
				}
			}
		}
	} else {
		err = qerr
	}
	rows.Close()
	return
}

// DelOverdueDevDate creat new device table
//删除最早的300条记录 ASC DESC
//DELETE FROM mytable WHERE timestamp IN (SELECT timestamp FROM mytable  ORDER BY timestamp ASC LIMIT 300)
//删除最早一天的记录
//DELETE FROM mytable WHERE timestamp < (SELECT (SELECT min(timestamp) FROM mytable) + INTERVAL '1 day')
//保留100天记录
//DELETE FROM mytable WHERE timestamp < (SELECT (SELECT max(timestamp) FROM mytable) - INTERVAL '100 day')
func (db *MYDB) DelOverdueDevDate(devid, days string) (err error) {
	if db == nil {
		return errors.New("mysql is not connect or connect error")
	}
	devid = filterbyte(devid)
	//strdel := `DELETE FROM ` +
	//devid +
	//`WHERE timestamp IN (SELECT timestamp FROM ` +
	//devid +
	//`order BY timestamp ASC LIMIT 300)` +
	//`;`
	var strdel string
	if db.DbType == "mysql" {
		strdel = `DELETE FROM ` +
			devid +
			` WHERE timestamp < (SELECT (SELECT max(timestamp) FROM ` +
			devid +
			`) - INTERVAL '` +
			days +
			` day')` +
			`;`
	}
	if db.DbType == "sqlite3" {
		idays, _ := strconv.Atoi(days)
		interval := idays * int((24 * time.Hour).Seconds())
		strdel = `DELETE FROM ` +
			devid +
			` WHERE timestamp < (SELECT (SELECT max(timestamp) FROM ` +
			devid +
			`) - ` +
			fmt.Sprintf("%d", interval) +
			`);`
	}
	log.Debugf(strdel)
	_, err = db.Exec(strdel)
	//db.Exec("VACUUM;")
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
