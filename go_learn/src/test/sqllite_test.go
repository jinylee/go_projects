package test

import (
	"os"

	"fmt"

	"github.com/go-xorm/xorm"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

// db
var db *xorm.Engine

// mapping structure
type User struct {
	Id   int64
	Name string
}

type LoginInfo struct {
	Id        int64
	IP        string
	UserId    int64
	TimeStamp string `xorm:"<-"`
	Nonuse    int    `xorm:"<-"`
}

func TestDB(t *testing.T) {

	// DB File Remove
	fmt.Println("============== DB File Remove ==============")
	f := "./test.db"
	os.Remove(f)

	// DB Open
	fmt.Println("============== DB Open ==============")
	var err error
	db, err = xorm.NewEngine("sqlite3", f)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	db.ShowSQL(true)

	// Create tables
	fmt.Println("============== Create Tables ==============")
	err = db.CreateTables(&User{}, &LoginInfo{})
	if err != nil {
		log.Fatal(err)
	}
	tables, err := db.DBMetas()
	if err != nil {
		log.Fatal(err)
	}
	for _, table := range tables {
		log.Info(table.Name)
	}

	// Insert
	fmt.Println("============== Insert Into Tables ==============")
	_, err = db.Insert(&User{1, "xlw"}, &User{2, "qew"}, &LoginInfo{1, "127.0.0.1", 1, "", 23})
	if err != nil {
		log.Fatal(err)
	}

	// Select
	fmt.Println("============== Select From Tables ==============")
	info := LoginInfo{}
	_, err = db.ID(1).Get(&info)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)

	users := make([]User, 0)
	err = db.Find(&users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
}
