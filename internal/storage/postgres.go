package storage

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
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

func PgStartConnection() string {
	db = pg.Connect(&pg.Options{
		Addr:     Stcon.Addr,
		User:     Stcon.User,
		Password: Stcon.Password,
		Database: Stcon.Database,
	})
	PgCreateTable()
	return "Postgres connected"
}

// miwatchsleep =============================================================
type MiwatchSleepRow struct {
	ID             uint64
	DateToday      time.Time
	StartPeriodInt int
	StartPeriod    time.Time
	EndPeriodInt   int
	EndPeriod      time.Time
	Dreams         string
	PeriodDuration int
} //============================

// miwatchhr ============================================================//Пульс;Метка времени;Дата;Время=
type MiwatchHrRow struct {
	ID           uint64
	DateToday    time.Time
	EventTimeInt int
	EventTime    time.Time
	Heartrate    int
} //=============================

func NewMiwatchSleepRow() *MiwatchSleepRow {
	return &MiwatchSleepRow{}
}
func NewMiwatchHrRow() *MiwatchHrRow {
	return &MiwatchHrRow{}
}
func PgCreateTable() {
	// miwatch sleep csv
	var pglink MiwatchSleepRow
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

func (row *MiwatchSleepRow) AddIfNotExist() {
	log.Println("PG toucghed", row.StartPeriodInt)
	// _, err := db.Model(row).
	// 	Where("start_period_int = ?", row.StartPeriodInt).
	// 	OnConflict("DO NOTHING"). // optional
	// 	SelectOrInsert()
}
func (row *MiwatchHrRow) AddIfNotExist() {
	log.Printf("Date: %d, HR: %d", row.EventTimeInt, row.Heartrate)
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
