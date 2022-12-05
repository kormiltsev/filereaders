package storage

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var db *pg.DB

func PgStartConnection() string {
	db = pg.Connect(&pg.Options{
		Addr:     DBcon.Adr,
		User:     DBcon.User,
		Password: DBcon.Password,
		Database: DBcon.DB,
	})
	PgCreateTable()
	return "Postgres connected"
}

func PgCreateTable() {
	var pglink PostgresLink
	err := db.CreateTable(&pglink, &orm.CreateTableOptions{
		Temp:          false,
		IfNotExists:   true,
		FKConstraints: true,
	})
	panicIf(err)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func PgQueryPOSTorSelect(row *PostgresLink) {
	_, err := db.Model(row).
		Where("original = ?", row.Original).WhereOr("alias = ?", row.Alias).
		OnConflict("DO NOTHING").
		SelectOrInsert()
	row.Err = err
}

func PgQueryPostAnyway(row *PostgresLink) error {
	_, err := db.Model(row).Insert()
	return err
}

func PgDBClose() {
	db.Close()
}
