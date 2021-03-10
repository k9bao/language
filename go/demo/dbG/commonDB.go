package dbg

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/mattn/go-sqlite3"
)

type CommonDB struct {
	Path string

	db   *sql.DB
	lock *sync.Mutex
}

func NewCommonDB(absFile string) (*CommonDB, error) {
	dir, _ := path.Split(absFile)
	if dir != "" {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	if _, err := os.Stat(absFile); os.IsNotExist(err) {
		file, err := os.Create(absFile)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := sql.Open("sqlite3-fast", absFile)
	if err != nil {
		return nil, err
	}

	return &CommonDB{Path: absFile, db: db, lock: &sync.Mutex{}}, nil
}

func (s *CommonDB) CreateTable(name string, items map[string]string) error {
	cmd := fmt.Sprintf("select count(*)  from sqlite_master where type='table' and name = '%v';", name)
	rows := s.db.QueryRow(cmd)
	var count int
	if err := rows.Scan(&count); err == nil {
		if count == 1 {
			s.db.Exec(fmt.Sprintf("DROP TABLE %v", name))
		}
	}
	cmd = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (id INTEGER NOT NULL PRIMARY KEY);", name)
	if _, err := s.db.Exec(cmd); err != nil {
		return err
	}
	//alter table tableName add columnName varchar(30) //增加列
	//alter table tableName drop column columnName //删除列
	//alter table tableName alter column columnName varchar(4000) //修改列类型
	//EXEC sp_rename 'tableName.column1' , 'column2' //修改列名称(column1->column2)
	for k, v := range items {
		cmd := fmt.Sprintf("ALTER TABLE %v ADD %v %v;", name, k, v)
		if _, err := s.db.Exec(cmd); err != nil {
			return err
		}
	}
	return nil
}

func (s *CommonDB) InsertData(name string, items map[string]string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	cols := ""
	rows := ""
	for k, v := range items {
		cols += k
		cols += ","

		rows += "'"
		rows += v
		rows += "'"
		rows += ","
	}
	cmd := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v);", name, cols[0:len(cols)-1], rows[0:len(rows)-1])
	_, err := s.db.Exec(cmd)
	return err
}

func (s *CommonDB) Release() error {
	s.db.Close()
	return nil
}

func InitSqlDriver() {
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
