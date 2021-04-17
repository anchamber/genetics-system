package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/anchamber/genetics-system/db/model"
)

type SytemDBMock struct {
	DB *sql.DB
}

var systems []*model.System = []*model.System{
	{Name: "doctor", Location: "tardis", Type: model.Techniplast, CleaningInterval: 90, LastCleaned: time.Now()},
	{Name: "rick", Location: "c-137", Type: model.Techniplast, CleaningInterval: 90, LastCleaned: time.Now()},
	{Name: "morty", Location: "herry-herpson", Type: model.Techniplast, CleaningInterval: 90, LastCleaned: time.Now()},
	{Name: "obi", Location: "high_ground", Type: model.Techniplast, CleaningInterval: 90, LastCleaned: time.Now()},
}

func NewMockDB() SytemDBMock {
	mock := SytemDBMock{DB: initDB("test.sqlite")}

	// for _, system := range systems {
	// 	mock.Insert(system)
	// }

	return mock
}

func (systemDB SytemDBMock) SelectAll() ([]*model.System, error) {
	selectStatement := `
		SELECT name, location, type, cleaning_interval, last_cleaned FROM systems;
	`
	rows, err := systemDB.DB.Query(selectStatement)
	if err != nil {
		log.Fatalf(`failed to select all`)
		return nil, err
	}

	var data []*model.System
	for rows.Next() {
		var entry model.System
		err = rows.Scan(&entry.Name, &entry.Location, &entry.Type, &entry.CleaningInterval, &entry.LastCleaned)
		if err != nil {
			return nil, err
		}
		data = append(data, &entry)
	}

	return data, nil
}

func (systemDB SytemDBMock) SelectByName(name string) (*model.System, error) {
	selectStatement := `
		SELECT name, location, type, cleaning_interval, last_cleaned 
		FROM systems
		WHERE name = $1;
	`
	rows, err := systemDB.DB.Query(selectStatement, name)
	if err != nil {
		log.Fatalf(`failed to select all`)
		return nil, err
	}

	var entry model.System
	if !rows.Next() {
		return nil, nil
	}
	err = rows.Scan(&entry.Name, &entry.Location, &entry.Type, &entry.CleaningInterval, &entry.LastCleaned)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (systemDB SytemDBMock) Insert(system *model.System) error {
	insertStatement := `
		INSERT INTO systems (name, location, type, cleaning_interval, last_cleaned)
			VALUES (?, ?, ?, ?, ?);
	`
	tx, err := systemDB.DB.Begin()
	if err != nil {
		fmt.Printf("failed to begin transaction\n")
		return err
	}

	statement, err := tx.Prepare(insertStatement)
	if err != nil {
		fmt.Printf("failed to prepare statement\n")
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(system.Name, system.Location, system.Type, system.CleaningInterval, system.LastCleaned)
	if err != nil {
		fmt.Printf("failed to execute statement\n")
		return err
	}
	tx.Commit()
	return nil
}

func (systemDB SytemDBMock) Update(system *model.System) error {
	insertStatement := `
		UPDATE systems 
			SET name = $1, location = $2 type = $3, cleaning_interval = $4, last_cleaned = $5
			WHERE name = $1;
	`
	tx, err := systemDB.DB.Begin()
	if err != nil {
		fmt.Printf("failed to begin transaction\n")
		return err
	}

	statement, err := tx.Prepare(insertStatement)
	if err != nil {
		fmt.Printf("failed to prepare statement\n")
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(system.Name, system.Location, system.Type, system.CleaningInterval, system.LastCleaned)
	if err != nil {
		fmt.Printf("failed to execute statement\n")
		return err
	}
	tx.Commit()
	return nil
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}

	CreateTables(db)
	CreateIndexes(db)

	return db
}

func CreateTables(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db\n")
		return err
	}

	systemTable := `
		CREATE TABLE IF NOT EXISTS systems(
			name							TEXT		PRIMARY KEY NOT NULL,
			location					TEXT,
			type							TEXT,
			cleaning_interval INT,
			last_cleaned			DATE
		);
	`

	_, err := db.Exec(systemTable)
	if err != nil {
		log.Fatalf("failed to create table\n")
		return err
	}
	return nil
}

func CreateIndexes(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db\n")
		return err
	}

	systemIndex := `
		CREATE UNIQUE INDEX IF NOT EXISTS idx_system_name ON systems(name);
	`

	_, err := db.Exec(systemIndex)
	if err != nil {
		log.Fatalf("failed to create table\n")
		return err
	}
	return nil
}
