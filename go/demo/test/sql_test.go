package test

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"

	"github.com/mattn/go-sqlite3"
)

func Init() {
	driver := &sqlite3.SQLiteDriver{
		ConnectHook: func(conn *sqlite3.SQLiteConn) (err error) {
			_, err = conn.Exec(`
				PRAGMA cache_size=64000;
				PRAGMA synchronous=NORMAL;
				PRAGMA journal_mode=WAL;
			`, nil)
			return
		},
	}
	sql.Register("sqlite3-fast", driver)
}

func CreateDB(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3-fast", name)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS photo (
		id      INTEGER NOT NULL PRIMARY KEY,
		feature BLOB,
		tag     TEXT
	);`); err != nil {
		return nil, err
	}
	return db, err
}

//INSERT INTO table (a,b,c) VALUES (1,2,3) ON DUPLICATE KEY UPDATE c=c+1;   如果原来存在，删除后重新插入 3.24.0 or greater support
//INSERT INTO table (a,b,c) VALUES (1,2,3),(4,5,6) ON DUPLICATE KEY UPDATE b=VALUES(b);   如果原来存在，更新。否则插入  3.24.0 or greater support
func insertDB(db *sql.DB) {
	//var inserter *sql.Stmt
	inserter, err := db.Prepare("INSERT INTO photo (id, feature, tag) VALUES (?, ?, ?)")
	if err != nil {
		panic("db.Prepare err")
	}
	for i := 0; i < 1000; i++ {
		if !existDB(db, i) {
			_, err = inserter.Exec(i, []byte(fmt.Sprintf("feature%d", i)), fmt.Sprintf("tagstr%d", i))
			if err != nil {
				panic("inserter.Exec err")
			}
		}
	}
}

func existDB(db *sql.DB, id int) bool {
	if err := db.QueryRow("SELECT id, tag FROM photo WHERE id = ?", id).Scan(&id); err == sql.ErrNoRows {
		return false
	}
	return true
}

func selectDB(db *sql.DB, id int, limit int) error {
	rows, err := db.Query("SELECT id, tag FROM photo WHERE id > ? limit ?", id, limit)
	if err != nil {
		panic("select err")
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var tag string
		if err := rows.Scan(&id, &tag); err != nil {
			return err
		}
		fmt.Println(id, tag)
	}
	return nil
}

func updateDB(db *sql.DB, id int) error {
	if existDB(db, id) {
		r, err := db.Exec("UPDATE photo SET tag=? WHERE id=?", "tagtest", id)
		if err != nil {
			panic("UPDATE err")
			return err
		}
		if rows, err := r.RowsAffected(); err != nil {
			panic("r.RowsAffected() err")
			return err
		} else if rows == 0 {
			panic("r.RowsAffected() == 0")
			return nil
		}
	}

	return nil
}

func deleteDB(db *sql.DB, ids []int) error {
	q := "DELETE FROM photo WHERE id IN (?" + strings.Repeat(",?", len(ids)-1) + ")"
	args := make([]interface{}, 0, len(ids))
	for _, id := range ids {
		args = append(args, id)
	}
	if r, err := db.Exec(q, args...); err != nil {
		panic("Exec delete err")
		return err
	} else {
		n, err := r.RowsAffected()
		if err != nil {
			panic(err)
			return err
		}
		if n == 0 {
			//panic("RowsAffected size is 0")
			fmt.Println("ids not exist")
		}
	}
	return nil
}

func TestSql(t *testing.T) {
	libVersion, libVersionNumber, sourceID := sqlite3.Version()
	fmt.Println(libVersion, "-", libVersionNumber, "-", sourceID)
	Init()
	db, err := CreateDB("test.db")
	if err != nil {
		panic("open err")
	}
	defer db.Close()

	insertDB(db)
	ids := []int{2, 3, 5}
	deleteDB(db, ids)
	updateDB(db, 6)
	selectDB(db, 9, 5)
	fmt.Println("end")
}
