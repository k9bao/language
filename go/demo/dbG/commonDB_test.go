package dbg

import (
	"log"
	"testing"
)

func TestDemoDB(t *testing.T) {
	log.Println("in.")
	InitSqlDriver()
	db, err := NewCommonDB("test.db")
	if err == nil {
		log.Println("NewCommonDB Success.", err)
	} else {
		log.Println("NewCommonDB fail.", err)
		return
	}

	cols := make(map[string]string)
	cols["col1"] = "INTEGER"
	cols["col2"] = "TEXT"
	cols["col3"] = "BOOL"
	if err = db.CreateTable("table1", cols); err == nil {
		log.Println("CreateTable Success.", err)
	} else {
		log.Println("CreateTable fail.", err)
		return
	}

	rows := make(map[string]string)
	rows["col1"] = "55"
	rows["col2"] = "testTEXT"
	rows["col3"] = "1" //bool 1:true,other:false
	if err = db.InsertData("table1", rows); err == nil {
		log.Println("InsertData Success.", err)
	} else {
		log.Println("InsertData fail.", err)
		return
	}
	log.Println("over.")
}
