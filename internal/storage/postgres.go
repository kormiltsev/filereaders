package storage

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sort"
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

type PGbase struct {
	MiwatchSleep []MiwatchSleepRow
	MiwatchHrRow []MiwatchHrRow
}

var Catalog = PGbase{
	MiwatchSleep: make([]MiwatchSleepRow, 0, 100),
	MiwatchHrRow: make([]MiwatchHrRow, 0, 100),
}

func (row *MiwatchSleepRow) AddIfNotExist() {

	Catalog.MiwatchSleep = append(Catalog.MiwatchSleep, *row)
	//log.Println("PG toucghed", row.StartPeriodInt)
	// _, err := db.Model(row).
	// 	Where("start_period_int = ?", row.StartPeriodInt).
	// 	OnConflict("DO NOTHING"). // optional
	// 	SelectOrInsert()
}
func (row *MiwatchHrRow) AddIfNotExist() {
	//log.Printf("Date: %d, HR: %d", row.EventTimeInt, row.Heartrate)
	Catalog.MiwatchHrRow = append(Catalog.MiwatchHrRow, *row)

}

func ConnectDB() error {
	file := "data/Catalog.json"
	f, err := os.OpenFile(file, os.O_CREATE|os.O_RDONLY, 0777)
	if err != nil {
		log.Println("not open file ", file, err)
		return err
	}
	err = json.NewDecoder(f).Decode(&Catalog)
	if err == io.EOF {
		log.Println("New Catalog.json created")
		return nil
	}
	if err != nil {
		log.Println("cand decode from json", err)
		return err
	}
	return nil
}

func CloseDB() {
	sort.Slice(Catalog.MiwatchHrRow, func(i, j int) bool {
		return Catalog.MiwatchHrRow[i].EventTimeInt > Catalog.MiwatchHrRow[j].EventTimeInt
	})
	sort.Slice(Catalog.MiwatchSleep, func(i, j int) bool {
		return Catalog.MiwatchSleep[i].StartPeriodInt > Catalog.MiwatchSleep[j].StartPeriodInt
	})

	file := "data/Catalog.json"
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0777) //os.O_APPEND|
	if err != nil {
		log.Println("not open file ", file, err)
		return
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	enc.Encode(&Catalog)
	if err != nil {
		log.Println("cand encode to json", err)
		return
	}
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

func SendCatalog(n int) *PGbase {

	return &PGbase{
		MiwatchSleep: Catalog.MiwatchSleep[:n],
		MiwatchHrRow: Catalog.MiwatchHrRow[:n],
	}
}
