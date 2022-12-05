package storage

import (
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/kormiltsev/filereaders/internal/app"
)

var db *pg.DB

type PGconfig struct {
	Addr     string
	User     string
	Password string
	Database string
}

var Stcon = PGconfig{
	Addr:     ":5432",
	User:     "postgres",
	Password: "root",
	Database: "postgres",
}

func PGconfigs() *PGconfig {
	return &Stcon
}
func PgStartConnection(set *app.Settingos) string {
	db = pg.Connect(&pg.Options{
		Addr:     set.Adress,
		User:     set.User,
		Password: set.Password,
		Database: set.DB,
	})
	PgCreateTable()
	return "Postgres connected"
}

func PgCreateTable() {
	// miwatch sleep csv
	var pglink app.PGMiwatchSleep
	err := db.CreateTable(&pglink, &orm.CreateTableOptions{
		Temp:          false,
		IfNotExists:   true,
		FKConstraints: true,
	})
	panicIf(err)
	// next table
	// var pglink app.MiwatchSleepCSV
	// err := db.CreateTable(&pglink, &orm.CreateTableOptions{
	// 	Temp:          false,
	// 	IfNotExists:   true,
	// 	FKConstraints: true,
	// })
	// panicIf(err)

}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

type PGWriter interface {
	AddIfNotExist()
}

func (row *app.MiwatchSleepRow) AddIfNotExist() {
	log.Println("PG toucghed", row.StartPeriodInt)
	// _, err := db.Model(row).
	// 	Where("start_period_int = ?", row.StartPeriodInt).
	// 	OnConflict("DO NOTHING"). // optional
	// 	SelectOrInsert()
}

// func PgQueryPOSTorSelect(row *PostgresLink) {
// 	_, err := db.Model(row).
// 		Where("original = ?", row.Original).WhereOr("alias = ?", row.Alias).
// 		OnConflict("DO NOTHING").
// 		SelectOrInsert()
// 	row.Err = err
// }

// func PgQueryPostAnyway(row *PostgresLink) error {
// 	_, err := db.Model(row).Insert()
// 	return err
// }

// func PgDBClose() {
// 	db.Close()
// }
