package main

import (
	//"github.com/go-sql-driver/mysql"
	"gorm.io/driver/sqlite"
	"log"

	"github.com/leaper-one/bubblebox/core/liverank"
	"github.com/sirupsen/logrus"
	//"gorm.io/driver/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	write *gorm.DB
	read  *gorm.DB
}

func main() {
	//OpenDB("./liver.db")
	OpenMysql("root:sexy0756@tcp(192.168.1.123:3306)/data?charset=utf8&parseTime=True&loc=Local")
}

func OpenMysql(dsn string) *DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}
	db.AutoMigrate(&liverrank.BiliRank{})
	return &DB{write: db, read: db}
}

func OpenDB(path string) *DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}
	db.AutoMigrate(&liverrank.BiliRank{})
	return &DB{write: db, read: db}
}

func (db *DB) Update() *gorm.DB {
	return db.write
}

func (db *DB) View() *gorm.DB {
	return db.read
}

func (db *DB) Debug() *DB {
	return &DB{
		write: db.write.Debug(),
		read:  db.read.Debug(),
	}
}

func (db *DB) Begin() *DB {
	tx := db.write.Begin()
	return &DB{
		write: tx,
		read:  db.read,
	}
}

func (db *DB) Rollback() error {
	return db.write.Rollback().Error
}

func (db *DB) Commit() error {
	return db.write.Commit().Error
}

func (db *DB) RollbackUnlessCommitted() {
	if err := db.write.Rollback().Error; err != nil {
		logrus.WithError(err).Error("DB: RollbackUnlessCommitted")
	}
}

func (db *DB) Tx(fn func(tx *DB) error) error {
	tx := db.Begin()
	defer tx.RollbackUnlessCommitted()

	if err := fn(tx); err != nil {
		return err
	}

	return tx.Commit()
}
